package tui

import (
	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type BaseModel struct {
	help  helpModel
	board boardModel
}

// InitialModel creates a new model with the given game.
func InitialModel(game *mancala.Game) *BaseModel {
	return &BaseModel{
		help:  initialHelpModel(),
		board: initialBoardModel(game),
	}
}

// Init is a Bubble Tea method to initialize the model.
func (m BaseModel) Init() tea.Cmd {
	boardCmd := m.board.Init()

	return tea.Batch(boardCmd)
}

// Update is a Bubble Tea method to update the model based on the given message.
func (m BaseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd, helpCmd, boardCmd tea.Cmd

	if v, ok := msg.(tea.KeyMsg); ok && key.Matches(v, keys.Quit) {
		return m, tea.Quit
	}

	m.board, boardCmd = m.board.Update(msg)

	m.help, helpCmd = m.help.Update(msg)

	return m, tea.Batch(cmd, helpCmd, boardCmd)
}

// View is a Bubble Tea method to render the current model.
func (m BaseModel) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, m.board.View(), m.help.View())
}
