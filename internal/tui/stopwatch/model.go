package stopwatch

import (
	"time"

	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	stopwach stopwatch.Model
}

func InitialModel() Model {
	return Model{
		stopwach: stopwatch.NewWithInterval(time.Second),
	}
}

func (m Model) Init() tea.Cmd {
	return m.stopwach.Init()
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.stopwach, cmd = m.stopwach.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	return m.stopwach.View()
}
