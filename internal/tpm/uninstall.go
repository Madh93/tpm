package tpm

import (
	"fmt"
	"log"
	"os"

	"github.com/Madh93/tpm/internal/terraform"
	"github.com/spf13/viper"
)

func Uninstall(provider *terraform.Provider) (err error) {
	fmt.Printf("Uninstalling %s...\n", provider)

	var installationPath = provider.InstallationPath()
	if viper.GetBool("debug") {
		log.Printf("Provider should be located in '%s' directory\n", installationPath)
	}

	// Check provider already exists
	if _, err = os.Stat(installationPath); os.IsNotExist(err) {
		fmt.Printf("%s not found! Ignoring...\n", provider)
		return nil
	}

	// Remove provider
	err = os.RemoveAll(installationPath)
	if err != nil {
		return
	}

	fmt.Printf("%s has been uninstalled sucessfully!\n", provider)

	return nil
}
