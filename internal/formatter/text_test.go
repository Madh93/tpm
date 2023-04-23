package formatter_test

import (
	"testing"

	"github.com/Madh93/tpm/internal/formatter"
	"github.com/Madh93/tpm/internal/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTextFormatterFormat(t *testing.T) {
	tests := []struct {
		name      string
		providers []*terraform.Provider
		expected  string
	}{
		{
			name:      "no installed provider",
			providers: nil,
			expected:  "No packages found.\n",
		},
		{
			name:      "one installed provider",
			providers: []*terraform.Provider{terraform.NewProvider(terraform.NewProviderName("hashicorp", "http", "3.2.1"), "linux", "amd64")},
			expected:  "'hashicorp/http@3.2.1' (linux/amd64)\n",
		},
		{
			name: "multiple installed providers",
			providers: []*terraform.Provider{
				terraform.NewProvider(terraform.NewProviderName("cloudflare", "cloudflare", "4.4.0"), "windows", "amd64"),
				terraform.NewProvider(terraform.NewProviderName("digitalocean", "digitalocean", "2.28.0"), "darwin", "arm64"),
				terraform.NewProvider(terraform.NewProviderName("hashicorp", "aws", "4.64.0"), "linux", "amd64"),
			},
			expected: "'cloudflare/cloudflare@4.4.0' (windows/amd64)\n" +
				"'digitalocean/digitalocean@2.28.0' (darwin/arm64)\n" +
				"'hashicorp/aws@4.64.0' (linux/amd64)\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			textFormatter := &formatter.TextFormatter{}
			output, err := textFormatter.Format(tt.providers)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, string(output))
		})
	}
}
