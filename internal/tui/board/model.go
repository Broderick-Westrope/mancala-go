package board

import (
	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
	"github.com/Broderick-Westrope/mancala-go/internal/tui/keys"
	"github.com/Broderick-Westrope/mancala-go/internal/tui/stopwatch"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	game         *mancala.Game
	length       int
	cursor       int
	stopwatch    stopwatch.Model
	topBorder    string
	bottomBorder string
}

func InitialModel(game *mancala.Game) Model {
	m := Model{
		cursor:    0,
		stopwatch: stopwatch.InitialModel(),
		game:      game,
		length:    len(game.Side1.Pits),
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return m.stopwatch.Init()
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var stopwatchCmd tea.Cmd

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
			if m.game.Turn == mancala.Player1Turn {
				m.game.ExecuteMove((m.length - 1) - m.cursor)
			} else {
				m.game.ExecuteMove(m.cursor)
			}
		case key.Matches(msg, keys.Keys.Quit):
			return m, tea.Quit
		}
	}

	m.stopwatch, stopwatchCmd = m.stopwatch.Update(msg)

	return m, stopwatchCmd
}

func (m Model) View() string {
	return m.buildBoard()
}
