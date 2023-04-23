package cmd

import (
	"log"

	"github.com/Madh93/tpm/internal/tpm"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all installed providers",
	Run: func(cmd *cobra.Command, args []string) {
		err := tpm.List()
		if err != nil {
			log.Fatal("Error: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Local Flags
	listCmd.Flags().StringP("output", "o", "text", "output in text, json, csv or table format")
	listCmd.Flags().VisitAll(bindCustomFlag)
}
