package tests

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	. "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"net"
	"os"
	"os/exec"
	"reference-application/internal/domain/program"
	"reference-application/internal/domain/version"
	"reference-application/internal/infrastructure/repositories"
	"testing"
	"time"
)

const binPath = "../../cmd/server"

type IntegrationTest struct {
	TestWithDatabase
	Addr string
}

// PrepareDraftVersion creates a version in the database and returns it.
func (tdb TestWithDatabase) PrepareDraftVersion(_ *testing.T) version.Version {
	_version, _program := NewDraftVersion()
	versionRepository := repositories.NewVersionRepository(tdb.DB)
	programRepository := repositories.NewProgramRepository(tdb.DB)
	programRepository.Save(context.TODO(), _program)
	versionRepository.Save(context.TODO(), _version)
	return _version
}

// PrepareVersionOnReview creates a version in the database and returns it.
func (tdb TestWithDatabase) PrepareVersionOnReview(_ *testing.T) version.Version {
	_version, _program := NewOnReviewVersion()
	versionRepository := repositories.NewVersionRepository(tdb.DB)
	programRepository := repositories.NewProgramRepository(tdb.DB)
	programRepository.Save(context.TODO(), _program)
	versionRepository.Save(context.TODO(), _version)
	return _version
}

// PrepareDraftVersionReadyToReview creates a prepared to review draft version in the database and returns it.
func (tdb TestWithDatabase) PrepareDraftVersionReadyToReview(_ *testing.T) version.Version {
	_version, _program := NewPreparedToReviewVersion()
	versionRepository := repositories.NewVersionRepository(tdb.DB)
	programRepository := repositories.NewProgramRepository(tdb.DB)
	programRepository.Save(context.TODO(), _program)
	versionRepository.Save(context.TODO(), _version)
	return _version
}

// SavePrograms saves programs to the database.
func (tdb TestWithDatabase) SavePrograms(_ *testing.T, programs []program.Program) {
	programRepository := repositories.NewProgramRepository(tdb.DB)
	for _, _program := range programs {
		programRepository.Save(context.Background(), _program)
	}
}

// SaveVersions saves versions to the database.
func (tdb TestWithDatabase) SaveVersions(_ *testing.T, versions []version.Version) {
	versionRepository := repositories.NewVersionRepository(tdb.DB)
	for _, _version := range versions {
		versionRepository.Save(context.Background(), _version)
	}
}

// TerminateDatabase kills the db container.
func (tdb TestWithDatabase) TerminateDatabase(t *testing.T) {
	require.NoError(t, tdb.dbContainer.Terminate(context.Background()))
}

// PrepareIntegrationTest starts the server and returns the address of the server and a database connection.
// The server is killed when the test is finished.
func PrepareIntegrationTest(t *testing.T) IntegrationTest {
	t.Helper()
	dbTest := PrepareTestWithDatabase(t)
	ln, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	addr := ln.Addr().String()
	require.NoError(t, ln.Close())
	pipeReader, pipeWriter, err := os.Pipe()
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, pipeWriter.Close())
	})
	go func() {
		_, err := io.Copy(os.Stdout, pipeReader)
		require.NoError(t, err)
	}()
	cmd := exec.Command("go", "run", binPath, "-addr", addr, "-dsn", dbTest.DSN, "-db", "postgres")
	cmd.Stdout = pipeWriter
	cmd.Stderr = pipeWriter
	err = cmd.Start()
	require.NoError(t, err)
	waitForHTTPServer(addr)
	t.Cleanup(func() {
		require.NoError(t, cmd.Process.Kill())
	})
	return IntegrationTest{
		Addr:             "http://" + addr,
		TestWithDatabase: dbTest,
	}
}

type TestWithDatabase struct {
	DB          *gorm.DB
	DSN         string
	dbContainer testcontainers.Container
}

// PrepareTestWithDatabase starts a postgres container and returns a database connection.
// The container is killed when the test is finished.
// The database is cleaned up when the test is finished.
// Migration is run on the database.
func PrepareTestWithDatabase(t *testing.T) TestWithDatabase {
	t.Helper()

	postgresContainer := runPostgresContainer(t)
	t.Cleanup(func() {
		_ = postgresContainer.Terminate(context.Background())
	})
	dsn, err := postgresContainer.ConnectionString(context.Background(), "sslmode=disable", "application_name=test")
	require.NoError(t, err)
	db, err := gorm.Open(postgres.Open(dsn))
	require.NoError(t, err)
	t.Cleanup(func() {
		sqlDB, err := db.DB()
		require.NoError(t, err)
		require.NoError(t, sqlDB.Close())
	})
	err = db.AutoMigrate(repositories.ProgramModel{}, repositories.VersionModel{}) // TODO https://github.com/SmotrovaLilit/golang-reference-application/issues/12
	require.NoError(t, err)

	return TestWithDatabase{
		DB:          db,
		DSN:         dsn,
		dbContainer: postgresContainer,
	}
}

type TestWithMockedDatabase struct {
	sqlmock.Sqlmock
	DB *gorm.DB
}

func PrepareTestWithMockedDatabase(t *testing.T) TestWithMockedDatabase {
	stdDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: stdDB}), &gorm.Config{})
	require.NoError(t, err)

	return TestWithMockedDatabase{
		DB:      db,
		Sqlmock: mock,
	}
}

func runPostgresContainer(t *testing.T) *PostgresContainer {
	postgresContainer, err := RunContainer(context.Background(),
		testcontainers.WithImage("postgres:latest"),
		WithDatabase("test"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	require.NoError(t, err)
	return postgresContainer
}

func waitForHTTPServer(addr string) {
	for {
		_, err := net.Dial("tcp", addr)
		if err == nil {
			return
		}
		time.Sleep(time.Millisecond * 100)
	}
}
