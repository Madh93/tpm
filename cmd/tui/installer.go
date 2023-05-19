package tui

import (
	"github.com/Madh93/tpm/internal/terraform"
	"github.com/Madh93/tpm/internal/tpm"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
)

type InstallRunner struct{}

func (r *InstallRunner) RunCmd(job ProviderJob) tea.Cmd {
	return func() tea.Msg {
		job.err = tpm.Install(job.provider, viper.GetBool("force"))
		return processedJobMsg(job)
	}
}

func (r InstallRunner) String() string {
	return "Installing"
}

func RunInstaller(providers []*terraform.Provider) (err error) {
	model := NewModel(&InstallRunner{}, providers)
	if _, err = tea.NewProgram(model).Run(); err != nil {
		return
	}
	return nil
}
