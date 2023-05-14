package tui

import (
	"fmt"

	"github.com/Madh93/tpm/internal/terraform"
	"github.com/Madh93/tpm/internal/tpm"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type uninstaller struct {
	spinner   spinner.Model
	providers []*terraform.Provider
	index     int
	err       error
}

func NewUninstallerModel(providers []*terraform.Provider) uninstaller {
	return uninstaller{
		spinner:   DefaultSpinner,
		providers: providers,
		index:     0,
	}
}

func RunUninstaller(providers []*terraform.Provider) (err error) {
	if _, err = tea.NewProgram(NewUninstallerModel(providers)).Run(); err != nil {
		return
	}
	return nil
}

func (u uninstaller) Init() tea.Cmd {
	return tea.Batch(
		uninstallProvider(u.providers[u.index]),
		u.spinner.Tick,
	)
}

func (u uninstaller) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return u, tea.Quit
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		u.spinner, cmd = u.spinner.Update(msg)
		return u, cmd

	case providerMsg:
		if u.index >= len(u.providers)-1 {
			return u, tea.Sequence(
				tea.Printf("%s %s", CheckMark, u.providers[u.index]),
				tea.Quit,
			)
		}
		u.index++
		return u, tea.Batch(
			tea.Printf("%s %s", CheckMark, u.providers[u.index-1]),
			uninstallProvider(u.providers[u.index]),
		)

	case errMsg:
		u.err = msg
		if u.index >= len(u.providers)-1 {
			return u, tea.Sequence(
				tea.Printf("%s %s Reason: %s", CrossMark, u.providers[u.index], u.err),
				tea.Quit,
			)
		}
		u.index++
		return u, tea.Batch(
			tea.Printf("%s %s Reason: %s", CrossMark, u.providers[u.index-1], u.err),
			uninstallProvider(u.providers[u.index]),
		)
	}

	return u, nil
}

func (u uninstaller) View() string {
	return fmt.Sprintf("%s Unistalling %s", u.spinner.View(), u.providers[u.index])
}

func uninstallProvider(provider *terraform.Provider) tea.Cmd {
	return func() tea.Msg {
		err := tpm.Uninstall(provider)
		if err != nil {
			return errMsg{err}
		}
		return providerMsg(provider.String())
	}
}
