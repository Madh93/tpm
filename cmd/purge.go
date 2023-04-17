package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/Madh93/tpm/internal/tpm"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:     "purge",
	Aliases: []string{"p"},
	Short:   "Purge ALL installed providers",
	Run: func(cmd *cobra.Command, args []string) {
		skipConfirmation := getBoolFlag(cmd, "yes")

		// Request user confirmation
		if !skipConfirmation {
			// Read user input
			log.Print("Are you sure you want to purge ALL installed providers? (yes/no)")
			var confirmation string
			_, err := fmt.Scanln(&confirmation)
			if err != nil {
				log.Fatal(err)
			}
			// Parse user input
			if strings.ToLower(confirmation) != "yes" && strings.ToLower(confirmation) != "y" {
				log.Println("Operation cancelled.")
				return
			}
		}

		// Purge ALL providers
		err := tpm.Purge()
		if err != nil {
			log.Fatal("Error: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)

	// Local Flags
	purgeCmd.Flags().BoolP("yes", "y", false, "skips the confirmation prompt and proceeds with the action")
}
