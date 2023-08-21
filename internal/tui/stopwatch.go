package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
)

type stopwatchModel struct {
	stopwach stopwatch.Model
}

func initialStopwatchModel() stopwatchModel {
	return stopwatchModel{
		stopwach: stopwatch.NewWithInterval(time.Second),
	}
}

// Init is a Bubble Tea method to initialize the model.
func (m stopwatchModel) Init() tea.Cmd {
	return m.stopwach.Init()
}

// Update is a Bubble Tea method to update the model based on the given message.
func (m stopwatchModel) Update(msg tea.Msg) (stopwatchModel, tea.Cmd) {
	var cmd tea.Cmd
	m.stopwach, cmd = m.stopwach.Update(msg)

	return m, cmd
}

// View is a Bubble Tea method to render the current model.
func (m stopwatchModel) View() string {
	return m.stopwach.View()
}
