package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

var (
	DefaultSpinner = spinner.New(spinner.WithSpinner(spinner.MiniDot))
	CheckMark      = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	CrossMark      = lipgloss.NewStyle().Foreground(lipgloss.Color("160")).SetString("⨯")
)

type providerMsg string

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }
