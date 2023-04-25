package parser

import (
	"runtime"

	"github.com/Madh93/tpm/internal/terraform"
	"gopkg.in/yaml.v3"
)

type YAMLProvidersFile struct {
	Providers []struct {
		Name string   `yaml:"name"`
		OS   []string `yaml:"os"`
		Arch []string `yaml:"arch"`
	} `yaml:"providers"`
}

type YAMLParser struct{}

func (f *YAMLParser) Parse(data []byte) (providers []*terraform.Provider, err error) {
	// Decode YAML
	var config YAMLProvidersFile
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	// Parse providers
	for _, provider := range config.Providers {
		osList := getListOrDefault(provider.OS, []string{runtime.GOOS})
		archList := getListOrDefault(provider.Arch, []string{runtime.GOARCH})
		for _, os := range osList {
			for _, arch := range archList {
				providerName, err := terraform.ParseProviderName(provider.Name)
				if err != nil {
					return nil, err
				}
				providers = append(providers, terraform.NewProvider(providerName, os, arch))
			}
		}
	}

	return providers, nil
}

func getListOrDefault(list, fallback []string) []string {
	if len(list) == 0 {
		return fallback
	}
	return list
}
