package version

import (
	"github.com/stretchr/testify/require"
	"reference-application/internal/domain/program"
	"testing"
)

func TestVersion_UpdateName(t *testing.T) {
	type fields struct {
		id        ID
		name      Name
		programID program.ID
		status    Status
	}
	type args struct {
		value Name
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
				value: MustNewName("new-name"),
			},
			wantErr: nil,
		},
		{
			name: "failed",
			fields: fields{
				id:        MustNewID("11a111cf-91f3-49dc-bb6d-ac4235635411"),
				name:      "name",
				programID: program.MustNewID("ecaffa6e-4302-4a46-ae72-44a7bd20dfd5"),
				status:    "NOT_DRAFT",
			},
			args: args{
				value: MustNewName("new-name"),
			},
			wantErr: ErrUpdateStatus,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Version{
				id:        tt.fields.id,
				name:      tt.fields.name,
				programID: tt.fields.programID,
				status:    tt.fields.status,
			}
			err := v.UpdateName(tt.args.value)
			require.ErrorIs(t, err, tt.wantErr)
			if err == nil {
				require.Equal(t, tt.args.value, v.name)
			} else {
				require.Equal(t, tt.fields.name, v.name)
			}
		})
	}
}
