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
