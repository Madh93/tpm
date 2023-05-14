package tui

import (
	"fmt"

	"github.com/Madh93/tpm/internal/terraform"
	"github.com/Madh93/tpm/internal/tpm"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
)

type installer struct {
	spinner   spinner.Model
	providers []*terraform.Provider
	index     int
	err       error
}

func NewInstallerModel(providers []*terraform.Provider) installer {
	return installer{
		spinner:   DefaultSpinner,
		providers: providers,
		index:     0,
	}
}

func RunInstaller(providers []*terraform.Provider) (err error) {
	if _, err = tea.NewProgram(NewInstallerModel(providers)).Run(); err != nil {
		return
	}
	return nil
}

func (i installer) Init() tea.Cmd {
	return tea.Batch(
		installProvider(i.providers[i.index]),
		i.spinner.Tick,
	)
}

func (i installer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return i, tea.Quit
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		i.spinner, cmd = i.spinner.Update(msg)
		return i, cmd

	case providerMsg:
		if i.index >= len(i.providers)-1 {
			return i, tea.Sequence(
				tea.Printf("%s %s", CheckMark, i.providers[i.index]),
				tea.Quit,
			)
		}
		i.index++
		return i, tea.Batch(
			tea.Printf("%s %s", CheckMark, i.providers[i.index-1]),
			installProvider(i.providers[i.index]),
		)

	case errMsg:
		i.err = msg
		if i.index >= len(i.providers)-1 {
			return i, tea.Sequence(
				tea.Printf("%s %s Reason: %s", CrossMark, i.providers[i.index], i.err),
				tea.Quit,
			)
		}
		i.index++
		return i, tea.Batch(
			tea.Printf("%s %s Reason: %s", CrossMark, i.providers[i.index-1], i.err),
			installProvider(i.providers[i.index]),
		)
	}

	return i, nil
}

func (i installer) View() string {
	return fmt.Sprintf("%s Installing %s", i.spinner.View(), i.providers[i.index])
}

func installProvider(provider *terraform.Provider) tea.Cmd {
	return func() tea.Msg {
		err := tpm.Install(provider, viper.GetBool("force"))
		if err != nil {
			return errMsg{err}
		}
		return providerMsg(provider.String())
	}
}
