package tui

import (
	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
	"github.com/Broderick-Westrope/mancala-go/internal/tui/board"
	"github.com/Broderick-Westrope/mancala-go/internal/tui/help"
	"github.com/Broderick-Westrope/mancala-go/internal/tui/keys"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	help  help.Model
	board board.Model
}

func InitialModel(game *mancala.Game) *model {
	return &model{
		help:  help.InitialModel(),
		board: board.InitialModel(game),
	}
}

func (m model) Init() tea.Cmd {
	helpCmd := m.help.Init()
	boardCmd := m.board.Init()

	return tea.Batch(helpCmd, boardCmd)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd, helpCmd, boardCmd tea.Cmd

	// Check for quit
	if v, ok := msg.(tea.KeyMsg); ok && key.Matches(v, keys.Keys.Quit) {
		return m, tea.Quit
	}

	m.board, boardCmd = m.board.Update(msg)

	m.help, helpCmd = m.help.Update(msg)

	return m, tea.Batch(cmd, helpCmd, boardCmd)
}

func (m model) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, m.board.View(), m.help.View())
}
