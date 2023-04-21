package terraform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProviderVersionString(t *testing.T) {
	tests := []struct {
		name     string
		version  ProviderVersion
		expected string
	}{
		{
			name:     "1.2.3",
			version:  ProviderVersion{Version: "1.2.3"},
			expected: "1.2.3",
		},
		{
			name:     "3.2.1",
			version:  ProviderVersion{Version: "3.2.1"},
			expected: "3.2.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.version.String())
		})
	}
}

func TestProviderVersionSemanticVersion(t *testing.T) {
	tests := []struct {
		name     string
		version  ProviderVersion
		expected string
		wantErr  bool
	}{
		{
			name:     "valid semantic version",
			version:  ProviderVersion{Version: "1.2.3"},
			expected: "1.2.3",
			wantErr:  false,
		},
		{
			name:     "invalid semantic version",
			version:  ProviderVersion{Version: "a.b.c"},
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			version, err := tt.version.SemanticVersion()

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, version)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, version)
				assert.Equal(t, tt.expected, version.String())
			}
		})
	}
}
