package tpm

import (
	"sort"

	"github.com/Madh93/tpm/internal/terraform"
)

var registry *terraform.Registry

func setLatestProviderVersion(provider *terraform.Provider) (err error) {
	versions, err := registry.GetVersions(provider)
	if err != nil {
		return
	}

	sort.Sort(versions)

	provider.SetVersion(versions.Last().String())

	return nil
}

func removeDuplicatedProviders(providers []*terraform.Provider) []*terraform.Provider {
	found := map[string]bool{}
	result := []*terraform.Provider{}

	for _, provider := range providers {
		if !found[provider.String()] {
			found[provider.String()] = true
			result = append(result, provider)
		}
	}

	return result
}
