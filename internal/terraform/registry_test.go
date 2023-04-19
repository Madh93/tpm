package terraform

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPackage(t *testing.T) {
	tests := []struct {
		name       string
		mockServer *httptest.Server
		provider   *Provider
		wantErr    bool
	}{
		{
			name:       "valid package",
			mockServer: makeValidServer(),
			provider:   NewProvider(NewProviderName("hashicorp", "null", "3.2.1"), "linux", "amd64"),
			wantErr:    false,
		},
		{
			name:       "package not found",
			mockServer: makeNotFoundServer(),
			provider:   NewProvider(NewProviderName("hashicorp", "whatever", "0.0.0"), "linux", "amd64"),
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := tt.mockServer
			defer server.Close()

			registry := Registry{baseURL: server.URL}
			pkg, err := registry.GetPackage(tt.provider)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, pkg)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, pkg.DownloadURL)
			}
		})
	}
}

func makeValidServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"protocols": [
				"5.0"
			],
			"os": "linux",
			"arch": "amd64",
			"filename": "terraform-provider-null_3.2.1_linux_amd64.zip",
			"download_url": "https://releases.hashicorp.com/terraform-provider-null/3.2.1/terraform-provider-null_3.2.1_linux_amd64.zip",
			"shasums_url": "https://releases.hashicorp.com/terraform-provider-null/3.2.1/terraform-provider-null_3.2.1_SHA256SUMS",
			"shasums_signature_url": "https://releases.hashicorp.com/terraform-provider-null/3.2.1/terraform-provider-null_3.2.1_SHA256SUMS.72D7468F.sig",
			"shasum": "74cb22c6700e48486b7cabefa10b33b801dfcab56f1a6ac9b6624531f3d36ea3",
			"signing_keys": {
				"gpg_public_keys": [
					{
						"key_id": "34365D9472D7468F",
						"ascii_armor": "-----BEGIN PGP PUBLIC KEY BLOCK-----",
						"trust_signature": "",
						"source": "HashiCorp",
						"source_url": "https://www.hashicorp.com/security.html"
					}
				]
			}
		}`))
	}))
}

func makeNotFoundServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{
			"errors": [
				"Not Found"
			]
		}`))
	}))
}
