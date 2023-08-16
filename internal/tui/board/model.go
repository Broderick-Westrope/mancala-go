package board

import (
	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
	"github.com/Broderick-Westrope/mancala-go/internal/tui/keys"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	game   *mancala.Game
	length int
	cursor int
}

func InitialModel(game *mancala.Game) Model {
	m := Model{
		cursor: 0,
	}
	m.UpdateGame(game)
	println(m.length)
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Left):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, keys.Keys.Right):
			if m.cursor < m.length-1 {
				m.cursor++
			}
		case key.Matches(msg, keys.Keys.Submit):
			m.game.ExecuteMove(m.cursor)
		case key.Matches(msg, keys.Keys.Quit):
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *Model) UpdateGame(game *mancala.Game) {
	m.game = game
	m.length = len(game.Side1.Pits)
}

func (m Model) View() string {
	return m.buildBoard()
}
