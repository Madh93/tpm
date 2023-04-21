package terraform

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

type Registry struct {
	baseURL string
}

func NewRegistry(hostname string) *Registry {
	return &Registry{
		baseURL: fmt.Sprintf("https://%s/v1/providers", hostname),
	}
}

func (r Registry) String() string {
	return fmt.Sprintf("'%s'", r.baseURL)
}

type GetVersionsResponse struct {
	Versions ProviderVersions `json:"versions"`
}

func (r *Registry) GetVersions(provider *Provider) (ProviderVersions, error) {
	url := fmt.Sprintf("%s/%s/%s/versions", r.baseURL, provider.Namespace(), provider.ProviderType())

	if viper.GetBool("debug") {
		log.Printf("Requesting the next url: '%s' \n", url)
	}

	// Get request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		break
	case 404:
		return nil, fmt.Errorf("provider not found")
	default:
		return nil, fmt.Errorf("unable to get provider versions: %q", resp.Status)
	}

	// Parse response
	var versionsResp GetVersionsResponse
	err = json.NewDecoder(resp.Body).Decode(&versionsResp)
	if err != nil {
		return nil, err
	}

	return versionsResp.Versions, nil
}

type GetPackageResponse struct {
	Protocols           []string `json:"protocols"`
	OS                  string   `json:"os"`
	Arch                string   `json:"arch"`
	Filename            string   `json:"filename"`
	DownloadURL         string   `json:"download_url"`
	SHASumsURL          string   `json:"shasums_url"`
	SHASumsSignatureURL string   `json:"shasums_signature_url"`
	SHASum              string   `json:"shasum"`
	SigningKeys         struct {
		GPGPublicKeys []struct {
			KeyID          string `json:"key_id"`
			ASCIIArmor     string `json:"ascii_armor"`
			TrustSignature string `json:"trust_signature"`
			Source         string `json:"source"`
			SourceURL      string `json:"source_url"`
		} `json:"gpg_public_keys"`
	} `json:"signing_keys"`
}

func (r *Registry) GetPackage(provider *Provider) (*GetPackageResponse, error) {
	url := fmt.Sprintf("%s/%s/%s/%s/download/%s/%s", r.baseURL, provider.Namespace(), provider.ProviderType(), provider.Version(), provider.OperatingSystem(), provider.Architecture())

	if viper.GetBool("debug") {
		log.Printf("Requesting the next url: '%s' \n", url)
	}

	// Get request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		break
	case 404:
		return nil, fmt.Errorf("provider not found")
	default:
		return nil, fmt.Errorf("unable to download provider: %q", resp.Status)
	}

	// Parse response
	var packageResp GetPackageResponse
	err = json.NewDecoder(resp.Body).Decode(&packageResp)
	if err != nil {
		return nil, err
	}

	return &packageResp, nil
}
