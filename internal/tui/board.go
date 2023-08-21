package tui

import (
	"log/slog"

	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type boardModel struct {
	game          *mancala.Game
	length        int
	cursor        int
	stopwatch     stopwatchModel
	topBorder     string
	bottomBorder  string
	isBotThinking bool
	spinner       spinner.Model
}

func initialBoardModel(game *mancala.Game) boardModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("5"))

	m := boardModel{
		cursor:    0,
		stopwatch: initialStopwatchModel(),
		game:      game,
		length:    len(game.Side1.Pits),
		spinner:   s,
	}

	return m
}

func (m boardModel) Init() tea.Cmd {
	return tea.Batch(m.stopwatch.Init(), m.spinner.Tick)
}

func (m boardModel) Update(msg tea.Msg) (boardModel, tea.Cmd) {
	var cmds []tea.Cmd
	var stopwatchCmd, spinnerCmd tea.Cmd

	m.spinner, spinnerCmd = m.spinner.Update(msg)
	m.stopwatch, stopwatchCmd = m.stopwatch.Update(msg)
	cmds = append(cmds, stopwatchCmd, spinnerCmd)

	if !m.isBotThinking {
		if p1, ok := m.game.Side1.Player.(*mancala.MinimaxBot); ok && m.game.Turn == mancala.Player1Turn {
			m.isBotThinking = true
			return m, getMinimaxMove(p1, m.game)
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Left):
			if m.cursor > 0 && !m.isBotThinking {
				m.cursor--
			}
		case key.Matches(msg, keys.Right):
			if m.cursor < m.length-1 && !m.isBotThinking {
				m.cursor++
			}
		case key.Matches(msg, keys.Submit):
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
		case key.Matches(msg, keys.Quit):
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

func (m boardModel) View() string {
	var message string

	if m.game.IsOver() {
		message = "Game over! "
	} else if m.isBotThinking {
		message = m.spinner.View() + "Bot is thinking | "
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

type notifyMsg struct{}

func notify() tea.Msg {
	return notifyMsg{}
}
