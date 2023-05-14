package cmd

import (
	"fmt"
	"log"
	"runtime"

	"github.com/Madh93/tpm/cmd/tui"
	"github.com/Madh93/tpm/internal/terraform"
	"github.com/Madh93/tpm/internal/tpm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var installCmd = &cobra.Command{
	Use:     "install [provider]",
	Aliases: []string{"i"},
	Short:   "Install a provider",
	Args: func(cmd *cobra.Command, args []string) error {
		installFromFile := getStringFlag(cmd, "from-file")
		if len(args) != 1 && installFromFile == "" {
			return fmt.Errorf("requires 1 arg when '--from-file' flag is not passed, received %d", len(args))
		}
		if len(args) >= 1 && installFromFile != "" {
			return fmt.Errorf("requires 0 arg when '--from-file' flag is passed, received %d", len(args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("force", getBoolFlag(cmd, "force"))
		var providers []*terraform.Provider
		var err error

		// Parse providers to install
		if getStringFlag(cmd, "from-file") != "" {
			providers, err = tpm.ParseProvidersFromFile(getStringFlag(cmd, "from-file"))
			if err != nil {
				log.Fatal("Error: ", err)
			}
		} else {
			for _, os := range getStringSliceFlag(cmd, "os") {
				for _, arch := range getStringSliceFlag(cmd, "arch") {
					providerName, err := terraform.ParseProviderName(args[0])
					if err != nil {
						log.Fatal("Error: ", err)
					}
					providers = append(providers, terraform.NewProvider(providerName, os, arch))
				}
			}
		}

		// Install providers
		err = tui.RunInstaller(providers)
		if err != nil {
			log.Fatal("Error: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Local Flags
	installCmd.Flags().Bool("force", false, "forces the installation of the provider even if it already exists")
	installCmd.Flags().StringP("from-file", "f", "", "installs providers defined in a 'providers.yml' file")
	installCmd.Flags().StringSliceP("os", "o", []string{runtime.GOOS}, "terraform provider operating system")
	installCmd.Flags().StringSliceP("arch", "a", []string{runtime.GOARCH}, "terraform provider architecture")
}
