package tui

import (
	"github.com/Broderick-Westrope/mancala-go/internal/tui/menu"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	menu menu.Model
}

func InitialModel() model {
	return model{
		menu: menu.InitialModel(),
	}
}

func (m model) Init() tea.Cmd {
	cmd := m.menu.Init()
	if cmd != nil {
		return cmd
	}

	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd, menuCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			return m, tea.Quit
		}
	}

	m.menu, menuCmd = m.menu.Update(msg)

	return m, tea.Batch(cmd, menuCmd)
}

func (m model) View() string {
	return m.menu.View()
}
