package help

import (
	"github.com/Broderick-Westrope/mancala-go/internal/tui/keys"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	help help.Model
	keys keys.KeyMap
}

func InitialModel() Model {
	return Model{
		help: help.New(),
		keys: keys.Keys,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}

	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	}

	return m, nil
}

func (m Model) View() string {
	return helpStyle.Render(m.help.View(m.keys))
}
