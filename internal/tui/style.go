package tui

import "github.com/charmbracelet/lipgloss"

var helpStyle = lipgloss.NewStyle().Height(2)

// Styles for the board
var (
	baseStyle      = lipgloss.NewStyle().Padding(0, 1)
	renderSelected = baseStyle.Copy().Foreground(lipgloss.Color("5")).Bold(true).Render
	renderDisabled = baseStyle.Copy().Foreground(lipgloss.Color("8")).Render
)
