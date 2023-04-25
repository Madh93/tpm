package tpm

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Madh93/tpm/internal/compression"
	"github.com/Madh93/tpm/internal/parser"
	"github.com/Madh93/tpm/internal/terraform"
	"github.com/spf13/viper"
)

func ParseProvidersFromFile(filename string) (providers []*terraform.Provider, err error) {
	if viper.GetBool("debug") {
		log.Printf("Reading '%s' providers file \n", filename)
	}

	// Setup input parser
	parser, err := parser.NewParser(filename)
	if err != nil {
		return
	}

	// Read file
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	// Parse providers
	providers, err = parser.Parse(data)
	if err != nil {
		return
	}

	return
}

func Install(provider *terraform.Provider, force bool) (err error) {
	fmt.Printf("Installing %s...\n", provider)

	// Setup registry
	registry = terraform.NewRegistry(viper.GetString("terraform_registry"))

	// Set latest version
	if provider.Version() == "latest" {
		err := setLatestProviderVersion(provider)
		if err != nil {
			return err
		}
	}

	// Check provider already exists
	if !force {
		if _, err = os.Stat(provider.InstallationPath()); !os.IsNotExist(err) {
			if viper.GetBool("debug") {
				log.Printf("Provider already installed in '%s' directory\n", provider.InstallationPath())
			}
			fmt.Printf("%s is already installed! Use '--force' to reinstall\n", provider)
			return nil
		}
	}

	// Download
	filename, err := downloadProvider(provider)
	if err != nil {
		return
	}

	// Extract
	err = extractProvider(provider, filename)
	if err != nil {
		return
	}

	fmt.Printf("%s has been installed sucessfully!\n", provider)

	return nil
}

func downloadProvider(provider *terraform.Provider) (filename string, err error) {
	fmt.Println("Downloading...")

	// Create Temporary file
	file, err := os.CreateTemp("", "")
	if err != nil {
		return
	}

	if viper.GetBool("debug") {
		log.Printf("Created tpm file under '%s' \n", file.Name())
	}

	// Get Download URL
	pkg, err := registry.GetPackage(provider)
	if err != nil {
		return
	}

	if viper.GetBool("debug") {
		log.Printf("Downloading provider from '%s'\n", pkg.DownloadURL)
	}

	// Download
	resp, err := http.Get(pkg.DownloadURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Save to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return
	}

	if viper.GetBool("debug") {
		log.Println("The download has finished sucessfully!")
	}

	return file.Name(), nil
}

func extractProvider(provider *terraform.Provider, filename string) (err error) {
	fmt.Println("Extracting...")

	destinationDir := provider.InstallationPath()

	if viper.GetBool("debug") {
		log.Printf("Provider will be extracted under '%s'\n", destinationDir)
	}

	// Extract
	err = compression.Unzip(filename, destinationDir)
	if err != nil {
		return
	}

	// Delete Temporary file
	err = os.Remove(filename)
	if err != nil {
		return
	}

	if viper.GetBool("debug") {
		log.Println("The provider has been extracted sucessfully!")
	}

	return nil
}
