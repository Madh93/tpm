package parser_test

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/Madh93/tpm/internal/parser"
	"github.com/Madh93/tpm/internal/terraform"
	"github.com/stretchr/testify/assert"
)

func TestYAMLParserParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []*terraform.Provider
	}{
		{
			name:     "no providers definition",
			input:    "",
			expected: nil,
		},
		{
			name: "one simple provider definition",
			input: `providers:
- name: hashicorp/http@3.2.1`,
			expected: []*terraform.Provider{terraform.NewProvider(terraform.NewProviderName("hashicorp", "http", "3.2.1"), runtime.GOOS, runtime.GOARCH)},
		},
		{
			name: "one provider definition with multiple arch and os",
			input: `providers:
- name: hashicorp/http@3.2.1
  os:
    - linux
    - darwin
  arch:
    - amd64
    - arm64`,
			expected: []*terraform.Provider{
				terraform.NewProvider(terraform.NewProviderName("hashicorp", "http", "3.2.1"), "linux", "amd64"),
				terraform.NewProvider(terraform.NewProviderName("hashicorp", "http", "3.2.1"), "linux", "arm64"),
				terraform.NewProvider(terraform.NewProviderName("hashicorp", "http", "3.2.1"), "darwin", "amd64"),
				terraform.NewProvider(terraform.NewProviderName("hashicorp", "http", "3.2.1"), "darwin", "arm64"),
			},
		},
		{
			name: "multiple providers definitions",
			input: `providers:
- name: cloudflare/cloudflare@4.4.0
- name: digitalocean/digitalocean@2.28.0
- name: hashicorp/aws@4.64.0`,
			expected: []*terraform.Provider{
				terraform.NewProvider(terraform.NewProviderName("cloudflare", "cloudflare", "4.4.0"), runtime.GOOS, runtime.GOARCH),
				terraform.NewProvider(terraform.NewProviderName("digitalocean", "digitalocean", "2.28.0"), runtime.GOOS, runtime.GOARCH),
				terraform.NewProvider(terraform.NewProviderName("hashicorp", "aws", "4.64.0"), runtime.GOOS, runtime.GOARCH),
			},
		},
		{
			name: "multiple providers definitions with multiple arch and os",
			input: `providers:
- name: cloudflare/cloudflare@4.4.0
  os:
    - darwin
  arch:
    - amd64
    - arm64
- name: digitalocean/digitalocean@2.28.0
  os:
    - windows
    - linux
  arch:
    - amd64
- name: hashicorp/aws@4.64.0
  os:
    - linux
    - darwin
  arch:
    - amd64
    - arm64`,
			expected: []*terraform.Provider{
				terraform.NewProvider(terraform.NewProviderName("cloudflare", "cloudflare", "4.4.0"), "darwin", "amd64"),
				terraform.NewProvider(terraform.NewProviderName("cloudflare", "cloudflare", "4.4.0"), "darwin", "arm64"),
				terraform.NewProvider(terraform.NewProviderName("digitalocean", "digitalocean", "2.28.0"), "windows", "amd64"),
				terraform.NewProvider(terraform.NewProviderName("digitalocean", "digitalocean", "2.28.0"), "linux", "amd64"),
				terraform.NewProvider(terraform.NewProviderName("hashicorp", "aws", "4.64.0"), "linux", "amd64"),
				terraform.NewProvider(terraform.NewProviderName("hashicorp", "aws", "4.64.0"), "linux", "arm64"),
				terraform.NewProvider(terraform.NewProviderName("hashicorp", "aws", "4.64.0"), "darwin", "amd64"),
				terraform.NewProvider(terraform.NewProviderName("hashicorp", "aws", "4.64.0"), "darwin", "arm64"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yamlParser := &parser.YAMLParser{}
			providers, err := yamlParser.Parse([]byte(tt.input))

			assert.NoError(t, err)
			assert.Equal(t, len(tt.expected), len(providers))
			if !reflect.DeepEqual(providers, tt.expected) {
				t.Errorf("TestYAMLParserParse(%q): expected provider %v, but got %v", tt.input, tt.expected, providers)
			}
		})
	}
}
