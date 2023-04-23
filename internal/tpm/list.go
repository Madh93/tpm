package tpm

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Madh93/tpm/internal/formatter"
	"github.com/Madh93/tpm/internal/pathutils"
	"github.com/Madh93/tpm/internal/terraform"
	"github.com/spf13/viper"
)

func List() (err error) {
	if viper.GetBool("debug") {
		log.Printf("Listing providers installed from '%s' registry\n", viper.GetString("terraform_registry"))
	}

	// Find providers
	providers, err := findProviders()
	if err != nil {
		return
	}

	// Setup output formatter
	formatter, err := formatter.NewFormatter(viper.GetString("output"))
	if err != nil {
		return
	}

	// Show output
	output, err := formatter.Format(providers)
	if err != nil {
		return
	}

	fmt.Print(string(output))

	return nil
}

func findProviders() (providers []*terraform.Provider, err error) {
	registryPath := filepath.Join(
		viper.GetString("terraform_plugin_cache_dir"),
		viper.GetString("terraform_registry"),
	)

	// Check registry path exists
	if _, err = os.Stat(registryPath); os.IsNotExist(err) {
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
