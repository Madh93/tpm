package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func bindCustomFlag(flag *pflag.Flag) {
	if flag.Name == "config" {
		return
	}
	name := strings.ReplaceAll(flag.Name, "-", "_")
	viper.BindPFlag(name, flag)
}

func getStringSliceFlag(cmd *cobra.Command, flag string) (value []string) {
	value, _ = cmd.Flags().GetStringSlice(flag)
	return
}

func getBoolFlag(cmd *cobra.Command, flag string) (value bool) {
	value, _ = cmd.Flags().GetBool(flag)
	return
}
