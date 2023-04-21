package terraform

import "fmt"

type ProviderPlatform struct {
	OS   string `json:"os"`
	Arch string `json:"arch"`
}

func (p ProviderPlatform) String() string {
	return fmt.Sprintf("%s/%s", p.OS, p.Arch)
}
