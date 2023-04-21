package terraform

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProviderPlatformString(t *testing.T) {
	tests := []struct {
		name     string
		platform ProviderPlatform
		expected string
	}{
		{
			name:     "linux amd64",
			platform: ProviderPlatform{OS: "linux", Arch: "amd64"},
			expected: "linux/amd64",
		},
		{
			name:     "darwin arm64",
			platform: ProviderPlatform{OS: "darwin", Arch: "arm64"},
			expected: "darwin/arm64",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.platform.String())
		})
	}

}
