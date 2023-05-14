package cmd

import (
	"log"
	"runtime"

	"github.com/Madh93/tpm/cmd/tui"
	"github.com/Madh93/tpm/internal/terraform"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:     "uninstall [provider]",
	Aliases: []string{"u"},
	Short:   "Uninstall a provider",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var providers []*terraform.Provider

		// Get providers to uninstall
		for _, os := range getStringSliceFlag(cmd, "os") {
			for _, arch := range getStringSliceFlag(cmd, "arch") {
				providerName, err := terraform.ParseProviderName(args[0])
				if err != nil {
					log.Fatal("Error: ", err)
				}
				providers = append(providers, terraform.NewProvider(providerName, os, arch))
			}
		}

		// Uninstall providers
		err := tui.RunUninstaller(providers)
		if err != nil {
			log.Fatal("Error: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	// Local Flags
	uninstallCmd.Flags().StringSliceP("os", "o", []string{runtime.GOOS}, "terraform provider operating system (empty to delete all architectures)")
	uninstallCmd.Flags().StringSliceP("arch", "a", []string{runtime.GOARCH}, "terraform provider architecture (empty to delete all operating systems)")
}
