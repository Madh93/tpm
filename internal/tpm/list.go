package tpm

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Madh93/tpm/internal/pathutils"
	"github.com/Madh93/tpm/internal/terraform"
	"github.com/spf13/viper"
)

func List() (err error) {
	fmt.Printf("Providers installed from '%s' registry: \n\n", viper.GetString("terraform_registry"))

	// Find providers
	providers, err := findProviders()
	if err != nil {
		return
	}

	// List providers
	for _, provider := range providers {
		fmt.Println(provider)
	}

	return nil
}

func findProviders() (providers []*terraform.Provider, err error) {
	registryPath := filepath.Join(
		viper.GetString("terraform_plugin_cache_dir"),
		viper.GetString("terraform_registry"),
	)

	// Check registry path exists
	if _, err = os.Stat(registryPath); os.IsNotExist(err) {
		fmt.Println("No packages found.")
		return nil, nil
	}

	registryDepth := pathutils.PathDepth(registryPath)

	// Find providers in registry path
	err = filepath.Walk(registryPath, func(path string, info os.FileInfo, errf error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && pathutils.PathDepth(path) == registryDepth+4 {
			provider, err := terraform.ParseProviderFromPath(path)
			if err != nil {
				return err
			}
			providers = append(providers, provider)
		}
		return nil
	})

	return providers, nil
}
