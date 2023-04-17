package terraform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProviderName_Parse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid input",
			input:   "namespace/type@1.0.0",
			wantErr: false,
		},
		{
			name:    "valid input with latest version",
			input:   "namespace/type@latest",
			wantErr: false,
		},
		{
			name:    "valid input without version",
			input:   "namespace/type",
			wantErr: false,
		},
		{
			name:    "invalid input",
			input:   "example",
			wantErr: true,
		},
		{
			name:    "invalid namespace",
			input:   "/type@1.0.0",
			wantErr: true,
		},
		{
			name:    "invalid multiple namespaces",
			input:   "namespace1/namespace2/type@1.0.0",
			wantErr: true,
		},
		{
			name:    "invalid name",
			input:   "namespace/@1.0.0",
			wantErr: true,
		},
		{
			name:    "invalid version",
			input:   "namespace/type@",
			wantErr: true,
		},
		{
			name:    "invalid multiple versions",
			input:   "namespace/type@1.0.0@2.0.0",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseProviderName(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
