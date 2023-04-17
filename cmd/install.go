package cmd

import (
	"log"
	"runtime"

	"github.com/Madh93/tpm/internal/terraform"
	"github.com/Madh93/tpm/internal/tpm"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install [provider]",
	Aliases: []string{"i"},
	Short:   "Install a provider",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var providers []*terraform.Provider

		// Get providers to install
		for _, os := range getStringSliceFlag(cmd, "os") {
			for _, arch := range getStringSliceFlag(cmd, "arch") {
				providerName, err := terraform.ParseProviderName(args[0])
				if err != nil {
					log.Fatal("Error: ", err)
				}
				providers = append(providers, terraform.NewProvider(providerName, os, arch))
			}
		}

		// Install providers
		for _, provider := range providers {
			err := tpm.Install(provider, getBoolFlag(cmd, "force"))
			if err != nil {
				log.Fatal("Error: ", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Local Flags
	installCmd.Flags().Bool("force", false, "forces the installation of the provider even if it already exists")
	installCmd.Flags().StringSliceP("os", "o", []string{runtime.GOOS}, "terraform provider operating system")
	installCmd.Flags().StringSliceP("arch", "a", []string{runtime.GOARCH}, "terraform provider architecture")
}
