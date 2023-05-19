package tui

import (
	"github.com/Madh93/tpm/internal/terraform"
	"github.com/Madh93/tpm/internal/tpm"
	tea "github.com/charmbracelet/bubbletea"
)

type UninstallRunner struct{}

func (r *UninstallRunner) RunCmd(job ProviderJob) tea.Cmd {
	return func() tea.Msg {
		job.err = tpm.Uninstall(job.provider)
		return processedJobMsg(job)
	}
}

func (r UninstallRunner) String() string {
	return "Uninstalling"
}

func RunUninstaller(providers []*terraform.Provider) (err error) {
	model := NewModel(&UninstallRunner{}, providers)
	if _, err = tea.NewProgram(model).Run(); err != nil {
		return
	}
	return nil
}
