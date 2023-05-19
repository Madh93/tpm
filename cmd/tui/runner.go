package tui

import (
	"fmt"

	"github.com/Madh93/tpm/internal/terraform"
	tea "github.com/charmbracelet/bubbletea"
)

type ProviderJob struct {
	id       int
	provider *terraform.Provider
	done     bool
	err      error
}

type JobRunner interface {
	RunCmd(job ProviderJob) tea.Cmd
	String() string
}

func NewRunner(runner string) (JobRunner, error) {
	switch runner {
	case "install":
		return &InstallRunner{}, nil
	case "uninstall":
		return &UninstallRunner{}, nil
	}
	return nil, fmt.Errorf("unsupported '%s' job runner", runner)
}
