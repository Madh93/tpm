package formatter_test

import (
	"testing"

	"github.com/Madh93/tpm/internal/formatter"
	"github.com/Madh93/tpm/internal/terraform"
	"github.com/stretchr/testify/assert"
)

func TestJSONFormatterFormat(t *testing.T) {
	tests := []struct {
		name      string
		providers []*terraform.Provider
		expected  string
	}{
		{
			name:      "no installed provider",
			providers: nil,
			expected:  "null",
		},
		{
			name:      "one installed provider",
			providers: []*terraform.Provider{terraform.NewProvider(terraform.NewProviderName("hashicorp", "http", "3.2.1"), "linux", "amd64")},
			expected: `[
  {
    "namespace": "hashicorp",
    "name": "http",
    "version": "3.2.1",
    "os": "linux",
    "arch": "amd64"
  }
]`,
		},
		{
			name: "multiple installed providers",
			providers: []*terraform.Provider{
				terraform.NewProvider(terraform.NewProviderName("cloudflare", "cloudflare", "4.4.0"), "windows", "amd64"),
				terraform.NewProvider(terraform.NewProviderName("digitalocean", "digitalocean", "2.28.0"), "darwin", "arm64"),
				terraform.NewProvider(terraform.NewProviderName("hashicorp", "aws", "4.64.0"), "linux", "amd64"),
			},
			expected: `[
  {
    "namespace": "cloudflare",
    "name": "cloudflare",
    "version": "4.4.0",
    "os": "windows",
    "arch": "amd64"
  },
  {
    "namespace": "digitalocean",
    "name": "digitalocean",
    "version": "2.28.0",
    "os": "darwin",
    "arch": "arm64"
  },
  {
    "namespace": "hashicorp",
    "name": "aws",
    "version": "4.64.0",
    "os": "linux",
    "arch": "amd64"
  }
]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonFormatter := &formatter.JSONFormatter{}
			output, err := jsonFormatter.Format(tt.providers)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, string(output))
		})
	}
}
