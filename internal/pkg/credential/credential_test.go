package credential

import "testing"

func TestValidation(t *testing.T) {
	tests := []struct {
		name    string
		cred    Credentials
		wantErr bool
	}{
		{
			"Ok",
			Credentials{
				ClientID:            "1",
				OrganizationID:      "2",
				SoftwareStatementID: "3",
			},
			false,
		},
		{
			"Empty Client ID",
			Credentials{
				OrganizationID:      "2",
				SoftwareStatementID: "3",
			},
			true,
		},
		{
			"Empty Organization ID",
			Credentials{
				ClientID:            "1",
				SoftwareStatementID: "3",
			},
			true,
		},
		{
			"Empty Software Statement ID",
			Credentials{
				ClientID:       "1",
				OrganizationID: "2",
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cred.Validate()

			if !tt.wantErr && got != nil {
				t.Error("unexpected error: ", got)
			}

			if tt.wantErr && got == nil {
				t.Error("expected error, got nil")
			}
		})
	}
}
