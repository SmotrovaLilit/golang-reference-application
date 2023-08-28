package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"reference-application/internal/infrastructure/repositories"
)

func ConnectDB(dbType, dsn string) *gorm.DB {
	var err error
	var db *gorm.DB

	switch dbType {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(repositories.ProgramModel{}, repositories.VersionModel{})
	if err != nil {
		panic("failed to migrate database")
	}

	return db
}
