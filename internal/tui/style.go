package tui

import "github.com/charmbracelet/lipgloss"

type colorFunc func(s ...string) string

func fg(color string) colorFunc {
	return lipgloss.NewStyle().Padding(0).Margin(0).Foreground(lipgloss.Color(color)).Render
}

var (
	Cyan    = fg("6")
	Faint   = fg("8")
	Magenta = fg("5")
	Red     = fg("1")
	NoStyle = lipgloss.NewStyle().Padding(0).Margin(0).Render
)
