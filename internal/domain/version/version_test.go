package version

import (
	"github.com/stretchr/testify/require"
	"reference-application/internal/domain/program"
	"reference-application/internal/pkg/optional"
	"testing"
)

func TestNewVersion(t *testing.T) {
	id := MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411")
	programID := program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5")
	name := MustNewName("name")
	got := NewVersion(id, name, programID)
	require.True(t, got.Status().IsDraft())
	require.Equal(t, id, got.ID())
	require.Equal(t, name, got.Name())
	require.Equal(t, programID, got.ProgramID())
	require.True(t, got.Description().IsEmpty())
	require.True(t, got.Number().IsEmpty())
}

func TestNewExistingVersion(t *testing.T) {
	id := MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411")
	programID := program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5")
	name := MustNewName("name")
	description := MustNewDescription("so long description")
	number := MustNewNumber("1.0.0")

	got := NewExistingVersion(
		id,
		name,
		programID,
		OnReviewStatus,
		optional.Of[Description](description),
		optional.Of[Number](number),
	)

	require.Equal(t, OnReviewStatus, got.Status())
	require.Equal(t, id, got.ID())
	require.Equal(t, name, got.Name())
	require.Equal(t, programID, got.ProgramID())
	require.True(t, got.Description().IsPresent())
	require.Equal(t, description, got.Description().Value())
	require.True(t, got.Number().IsPresent())
	require.Equal(t, number, got.Number().Value())
}

//nolint:funlen
func TestVersion_Update(t *testing.T) {
	type fields struct {
		id          ID
		name        Name
		description optional.Optional[Description]
		number      optional.Optional[Number]
		programID   program.ID
		status      Status
	}
	type args struct {
		description optional.Optional[Description]
		name        Name
		number      optional.Optional[Number]
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "success",
			fields: fields{
				id:        MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411"),
				name:      "name",
				number:    optional.Of[Number](MustNewNumber("1.0.0")),
				programID: program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5"),
				status:    DraftStatus,
			},
			args: args{
				description: optional.Of[Description](MustNewDescription("new-description")),
				name:        MustNewName("new-name"),
				number:      optional.Of[Number](MustNewNumber("1.0.1")),
			},
			wantErr: nil,
		},
		{
			name: "failed",
			fields: fields{
				id:        MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411"),
				name:      "name",
				number:    optional.Of[Number](MustNewNumber("1.0.0")),
				programID: program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5"),
				status:    OnReviewStatus,
			},
			args: args{
				description: optional.Of[Description](MustNewDescription("new-description")),
				name:        MustNewName("new-name"),
				number:      optional.Of[Number](MustNewNumber("1.0.1")),
			},
			wantErr: ErrUpdateVersionStatus,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Version{
				id:          tt.fields.id,
				name:        tt.fields.name,
				description: tt.fields.description,
				number:      tt.fields.number,
				programID:   tt.fields.programID,
				status:      tt.fields.status,
			}
			err := v.Update(
				tt.args.name,
				tt.args.description,
				tt.args.number,
			)
			require.ErrorIs(t, err, tt.wantErr)
			if err == nil {
				require.Equal(t, tt.args.description, v.description)
				require.Equal(t, tt.args.name, v.name)
				require.Equal(t, tt.args.number, v.number)
			} else {
				require.Equal(t, tt.fields.description, v.description)
				require.Equal(t, tt.fields.name, v.name)
				require.Equal(t, tt.fields.number, v.number)
			}
		})
	}
}

//nolint:funlen
func TestVersion_SendToReview(t *testing.T) {
	type fields struct {
		id          ID
		name        Name
		programID   program.ID
		status      Status
		description optional.Optional[Description]
		number      optional.Optional[Number]
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "success",
			fields: fields{
				id:          MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411"),
				name:        MustNewName("name"),
				number:      optional.Of[Number](MustNewNumber("1.0.0")),
				programID:   program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5"),
				status:      DraftStatus,
				description: optional.Of[Description](MustNewDescription("description")),
			},
		},
		{
			name: "status_failed",
			fields: fields{
				id:          MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411"),
				name:        MustNewName("name"),
				number:      optional.Of[Number](MustNewNumber("1.0.0")),
				programID:   program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5"),
				status:      OnReviewStatus,
				description: optional.Of[Description](MustNewDescription("description")),
			},
			wantErr: ErrInvalidStatusToSendToReview,
		},
		{
			name: "description_is_empty",
			fields: fields{
				id:          MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411"),
				name:        MustNewName("name"),
				number:      optional.Of[Number](MustNewNumber("1.0.0")),
				programID:   program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5"),
				status:      DraftStatus,
				description: optional.Empty[Description](),
			},
			wantErr: ErrEmptyDescription,
		},
		{
			name: "number_is_empty",
			fields: fields{
				id:          MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411"),
				name:        MustNewName("name"),
				number:      optional.Empty[Number](),
				programID:   program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5"),
				status:      DraftStatus,
				description: optional.Of[Description](MustNewDescription("description")),
			},
			wantErr: ErrEmptyNumber,
		},
		{
			name: "description_validation_for_review_failed",
			fields: fields{
				id:          MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411"),
				name:        MustNewName("name"),
				number:      optional.Of[Number](MustNewNumber("1.0.0")),
				programID:   program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5"),
				status:      DraftStatus,
				description: optional.Of[Description](MustNewDescription("sh")),
			},
			wantErr: ErrDescriptionLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Version{
				id:          tt.fields.id,
				name:        tt.fields.name,
				programID:   tt.fields.programID,
				status:      tt.fields.status,
				description: tt.fields.description,
				number:      tt.fields.number,
			}
			err := v.SendToReview()
			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, OnReviewStatus, v.status)
			} else {
				require.ErrorIs(t, err, tt.wantErr)
				require.Equal(t, tt.fields.status, v.status)
			}
		})
	}
}
