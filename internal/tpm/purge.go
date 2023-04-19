package tpm

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func Purge() (err error) {
	fmt.Println("Removing all providers...")

	registryPath := filepath.Join(
		viper.GetString("terraform_plugin_cache_dir"),
		viper.GetString("terraform_registry"),
	)

	if viper.GetBool("debug") {
		log.Printf("Providers should be located in '%s' directory\n", registryPath)
	}

	// Check provider already exists
	if _, err = os.Stat(registryPath); os.IsNotExist(err) {
		fmt.Printf("Registry path under '%s' does not exist! Ignoring...\n", registryPath)
		return nil
	}

	// Remove provider
	err = os.RemoveAll(registryPath)
	if err != nil {
		return
	}

	fmt.Println("All providers were removed sucessfully!")

	return nil
}
