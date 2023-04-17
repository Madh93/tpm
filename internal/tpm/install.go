package tpm

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Madh93/tpm/internal/compression"
	"github.com/Madh93/tpm/internal/terraform"
	"github.com/spf13/viper"
)

func Install(provider *terraform.Provider, force bool) (err error) {
	log.Printf("Installing %s...\n", provider)

	// Check provider already exists
	if !force {
		if _, err = os.Stat(provider.InstallationPath()); !os.IsNotExist(err) {
			log.Printf("Provider already exists in '%s' directory! Use '--force' to reinstall\n", provider.InstallationPath())
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

	log.Printf("%s has been installed sucessfully!\n", provider)

	return nil
}

func downloadProvider(provider *terraform.Provider) (filename string, err error) {
	log.Println("Downloading...")

	// Create Temporary file
	file, err := os.CreateTemp("", "")
	if err != nil {
		return
	}

	if viper.GetBool("debug") {
		log.Printf("Created tpm file under '%s' \n", file.Name())
	}

	// Get Download URL
	registry := terraform.NewRegistry(viper.GetString("terraform_registry"))
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
	log.Println("Extracting...")

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
