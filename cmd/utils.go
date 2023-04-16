package cmd

import (
	"strings"

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
