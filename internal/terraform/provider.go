package terraform

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

type ProviderName struct {
	namespace    string
	providerType string
	version      string
}

func NewProviderName(namespace, providerType, version string) *ProviderName {
	return &ProviderName{
		namespace:    strings.ToLower(namespace),
		providerType: strings.ToLower(providerType),
		version:      strings.ToLower(version),
	}
}

func ParseProviderName(s string) (*ProviderName, error) {
	parts := strings.Split(s, "/")
	typeWithVersion := parts[len(parts)-1]
	typeWithVersionParts := strings.Split(typeWithVersion, "@")
	if len(parts) != 2 || len(typeWithVersionParts) > 2 {
		return nil, errors.New("incorrect provider name format, expected '<namespace>/<type>[@<version>]'")
	}

	// Get Namespace
	namespace := strings.TrimSpace(parts[0])
	if namespace == "" {
		return nil, errors.New("invalid provider namespace, expected non-empty value")
	}

	// Get Type
	providerType := strings.TrimSpace(typeWithVersionParts[0])
	if providerType == "" {
		return nil, errors.New("invalid provider type, expected non-empty value")
	}

	// Get Version
	version := "latest"
	if len(typeWithVersionParts) > 1 {
		version = strings.TrimSpace(typeWithVersionParts[1])

		if version == "" {
			return nil, errors.New("invalid provider version, expected non-empty value")
		}
	}

	return NewProviderName(namespace, providerType, version), nil
}

type Provider struct {
	name            *ProviderName
	operatingSystem string
	architecture    string
}

func NewProvider(name *ProviderName, os, arch string) *Provider {
	return &Provider{
		name:            name,
		operatingSystem: strings.ToLower(os),
		architecture:    strings.ToLower(arch),
	}
}

func ParseProviderFromPath(path string) (*Provider, error) {
	// Get parts
	parts := strings.Split(path, string(os.PathSeparator))

	// Verify path format
	pattern := `^.*\/[^\/]+\/[^\/]+\/[^\/]+\/\d+\.\d+\.\d+\/[^_]+_[^\/]+$`
	match, err := regexp.MatchString(pattern, path)
	if err != nil {
		return nil, err
	}
	if !match && runtime.GOOS != "windows" {
		return nil, errors.New("invalid path format, expected something like '.../<namespace>/<type>/<version>/<os>_<arch>'")
	}

	// Get provider data
	os := strings.Split(parts[len(parts)-1], "_")[0]
	arch := strings.Split(parts[len(parts)-1], "_")[1]
	version := parts[len(parts)-2]
	providerType := parts[len(parts)-3]
	namespace := parts[len(parts)-4]

	providerName := NewProviderName(namespace, providerType, version)
	return NewProvider(providerName, os, arch), nil
}

func (p *Provider) Namespace() string {
	return p.name.namespace
}

func (p *Provider) ProviderType() string {
	return p.name.providerType
}

func (p *Provider) Version() string {
	return p.name.version
}

func (p *Provider) SetVersion(version string) {
	p.name.version = version
}

func (p *Provider) OperatingSystem() string {
	return p.operatingSystem
}

func (p *Provider) Architecture() string {
	return p.architecture
}

func (p Provider) String() string {
	return fmt.Sprintf("'%s/%s@%s' (%s/%s)", p.name.namespace, p.name.providerType, p.name.version, p.operatingSystem, p.architecture)
}

func (p Provider) InstallationPath() string {
	return filepath.Join(
		viper.GetString("terraform_plugin_cache_dir"),
		viper.GetString("terraform_registry"),
		p.name.namespace,
		p.name.providerType,
		p.name.version,
		fmt.Sprintf("%s_%s", p.operatingSystem, p.architecture),
	)
}
