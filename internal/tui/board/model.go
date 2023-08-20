package board

import (
	"log/slog"

	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
	"github.com/Broderick-Westrope/mancala-go/internal/tui/keys"
	"github.com/Broderick-Westrope/mancala-go/internal/tui/stopwatch"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	game          *mancala.Game
	length        int
	cursor        int
	stopwatch     stopwatch.Model
	topBorder     string
	bottomBorder  string
	isBotThinking bool
}

type NotifyMsg struct{}

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
	var cmds []tea.Cmd
	var stopwatchCmd tea.Cmd

	m.stopwatch, stopwatchCmd = m.stopwatch.Update(msg)
	cmds = append(cmds, stopwatchCmd)

	if !m.isBotThinking {
		if p1, ok := m.game.Side1.Player.(*mancala.MinimaxBot); ok && m.game.Turn == mancala.Player1Turn {
			m.isBotThinking = true
			return m, getMinimaxMove(p1, m.game)
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Left):
			if m.cursor > 0 && !m.isBotThinking {
				m.cursor--
			}
		case key.Matches(msg, keys.Keys.Right):
			if m.cursor < m.length-1 && !m.isBotThinking {
				m.cursor++
			}
		case key.Matches(msg, keys.Keys.Submit):
			if m.isBotThinking {
				break
			}
			cursor := m.cursor
			if m.game.Turn == mancala.Player1Turn {
				cursor = (m.length - 1) - cursor
			}
			err := m.game.ExecuteMove(cursor)
			if err != nil {
				slog.Error(err.Error())
			}
			cmds = append(cmds, notify)
		case key.Matches(msg, keys.Keys.Quit):
			return m, tea.Quit
		}
	case moveMsg:
		err := m.game.ExecuteMove(int(msg))
		if err != nil {
			slog.Error(err.Error())
			panic(err)
		}
		m.isBotThinking = false
		cmds = append(cmds, notify)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	var message string

	if m.game.IsOver() {
		message = "Game over! "
	} else if m.isBotThinking {
		message = "Bot is thinking | "
	}

	message += m.stopwatch.View()

	return m.buildBoard(message)
}

type moveMsg int

func getMinimaxMove(bot *mancala.MinimaxBot, game *mancala.Game) tea.Cmd {
	return func() tea.Msg {
		return moveMsg(bot.GetMove(game))
	}
}

func notify() tea.Msg {
	return NotifyMsg{}
}
