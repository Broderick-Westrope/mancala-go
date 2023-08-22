package tui

import "github.com/charmbracelet/lipgloss"

// Styles for the help text.
var (
	helpHeight = 2
	helpStyle  = lipgloss.NewStyle().Height(helpHeight)
)

// Styles for the board.
var (
	baseStyle      = lipgloss.NewStyle().Padding(0, 1)
	renderSelected = baseStyle.Copy().Foreground(lipgloss.Color("5")).Bold(true).Render
	renderDisabled = baseStyle.Copy().Foreground(lipgloss.Color("8")).Render
)
