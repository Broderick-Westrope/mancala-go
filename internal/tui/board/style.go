package board

import "github.com/charmbracelet/lipgloss"

var (
	baseStyle      = lipgloss.NewStyle().Padding(0, 1)
	renderSelected = baseStyle.Copy().Foreground(lipgloss.Color("3")).Render
	renderDisabled = baseStyle.Copy().Foreground(lipgloss.Color("8")).Render
)
