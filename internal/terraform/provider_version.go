package terraform

import "github.com/Masterminds/semver/v3"

type ProviderVersion struct {
	Version   string             `json:"version"`
	Protocols []string           `json:"protocols"`
	Platforms []ProviderPlatform `json:"platforms"`
}

func (p ProviderVersion) String() string {
	return p.Version
}

func (p ProviderVersion) SemanticVersion() (version *semver.Version, err error) {
	version, err = semver.NewVersion(p.Version)
	if err != nil {
		return
	}
	return
}
