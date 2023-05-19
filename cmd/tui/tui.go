package tui

import (
	"fmt"
	"time"

	"github.com/Madh93/tpm/internal/mathutils"
	"github.com/Madh93/tpm/internal/terraform"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

var (
	DefaultSpinner = spinner.New(spinner.WithSpinner(spinner.MiniDot))
	CheckMark      = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	CrossMark      = lipgloss.NewStyle().Foreground(lipgloss.Color("160")).SetString("⨯")
)

type processedJobMsg ProviderJob

type tearDownMsg bool

type model struct {
	spinner           spinner.Model
	runner            JobRunner
	jobs              []ProviderJob
	maxConcurrentJobs int
	doneJobs          int
	index             int
}

func NewModel(runner JobRunner, providers []*terraform.Provider) *model {
	// Setup jobs
	jobs := []ProviderJob{}
	for i, provider := range providers {
		jobs = append(jobs, ProviderJob{id: i, provider: provider, done: false, err: nil})
	}

	return &model{
		spinner:           DefaultSpinner,
		runner:            runner,
		jobs:              jobs,
		maxConcurrentJobs: viper.GetInt("jobs"),
	}
}

func (m *model) Init() tea.Cmd {
	// Setup spinner
	cmds := []tea.Cmd{m.spinner.Tick}

	// Add inital jobs
	for _, job := range m.jobs[:mathutils.Min(m.maxConcurrentJobs, len(m.jobs))] {
		cmds = append(cmds, m.runner.RunCmd(job))
		m.index++
	}

	return tea.Batch(cmds...)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case processedJobMsg:
		m.FinishJob(msg)
		output := m.GetJobOutput(msg)

		// Latest job! Quitting...
		if m.doneJobs >= len(m.jobs) {
			return m, tea.Sequence(
				tea.Println(output),
				tearDownCmd(),
			)
		}

		// Print result to output
		cmds := []tea.Cmd{tea.Println(output)}

		// Adding pending jobs
		if m.index <= len(m.jobs) {
			cmds = append(cmds, m.runner.RunCmd(m.jobs[m.index-1]))
		}

		return m, tea.Batch(cmds...)

	case tearDownMsg:
		for {
			time.Sleep(time.Millisecond * time.Duration(100))
			allDone := true
			for _, job := range m.jobs {
				if !job.done {
					allDone = false
					break
				}
			}

			if allDone {
				break
			}
		}
		return m, tea.Quit
	}

	return m, nil
}

func (m model) View() string {
	var view string
	var currentJobs int

	for _, job := range m.jobs {
		// Limit 'Installing' providers to the Max Concurrent Jobs
		if currentJobs >= m.maxConcurrentJobs {
			break
		}
		if !job.done {
			view += fmt.Sprintf("%s %s %s\n", m.spinner.View(), m.runner, job.provider)
			currentJobs++
		}
	}

	return view
}

func (m *model) FinishJob(job processedJobMsg) {
	m.jobs[job.id].done = true
	m.index++
	m.doneJobs++
}

func (m model) GetJobOutput(job processedJobMsg) string {
	if job.err != nil {
		return fmt.Sprintf("%s %s Reason: %s", CrossMark, job.provider, job.err)
	}

	return fmt.Sprintf("%s %s", CheckMark, job.provider)
}

func tearDownCmd() tea.Cmd {
	return func() tea.Msg {
		return tearDownMsg(true)
	}
}
