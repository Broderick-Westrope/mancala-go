package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type helpModel struct {
	help help.Model
	keys keyMap
}

func initialHelpModel() helpModel {
	return helpModel{
		help: help.New(),
		keys: keys,
	}
}

// Update is a Bubble Tea method to update the model based on the given message.
func (m helpModel) Update(msg tea.Msg) (helpModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Help) {
			m.help.ShowAll = !m.help.ShowAll
		}
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	}

	return m, nil
}

// View is a Bubble Tea method to render the current model.
func (m helpModel) View() string {
	return helpStyle.Render(m.help.View(m.keys))
}
