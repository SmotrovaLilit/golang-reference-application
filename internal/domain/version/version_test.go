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
}

func TestNewExistingVersion(t *testing.T) {
	id := MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411")
	programID := program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5")
	name := MustNewName("name")
	description := MustNewDescription("so long description")

	got := NewExistingVersion(id, name, programID, OnReviewStatus, optional.Of[Description](description))

	require.Equal(t, OnReviewStatus, got.Status())
	require.Equal(t, id, got.ID())
	require.Equal(t, name, got.Name())
	require.Equal(t, programID, got.ProgramID())
	require.True(t, got.Description().IsPresent())
	require.Equal(t, description, got.Description().Value())
}

func TestVersion_Update(t *testing.T) {
	type fields struct {
		id          ID
		name        Name
		description optional.Optional[Description]
		programID   program.ID
		status      Status
	}
	type args struct {
		description optional.Optional[Description]
		name        Name
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
				programID: program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5"),
				status:    DraftStatus,
			},
			args: args{
				description: optional.Of[Description](MustNewDescription("new-description")),
				name:        MustNewName("new-name"),
			},
			wantErr: nil,
		},
		{
			name: "failed",
			fields: fields{
				id:        MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411"),
				name:      "name",
				programID: program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5"),
				status:    OnReviewStatus,
			},
			args: args{
				description: optional.Of[Description](MustNewDescription("new-description")),
				name:        MustNewName("new-name"),
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
				programID:   tt.fields.programID,
				status:      tt.fields.status,
			}
			err := v.Update(tt.args.name, tt.args.description)
			require.ErrorIs(t, err, tt.wantErr)
			if err == nil {
				require.Equal(t, tt.args.description, v.description)
				require.Equal(t, tt.args.name, v.name)
			} else {
				require.Equal(t, tt.fields.description, v.description)
				require.Equal(t, tt.fields.name, v.name)
			}
		})
	}
}
