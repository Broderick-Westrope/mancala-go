package tui

import (
	"fmt"
	"os"
	"strings"

	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
	"github.com/Broderick-Westrope/mancala-go/internal/tui/board"
	"github.com/Broderick-Westrope/mancala-go/internal/tui/help"
	"github.com/Broderick-Westrope/mancala-go/internal/tui/keys"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	menuView = iota
	boardView
)

type model struct {
	menuOptions []string
	cursor      int
	title       string
	help        help.Model
	board       board.Model
	currentView uint
}

func InitialModel(game *mancala.Game) model {
	result, err := os.ReadFile("assets/mancala_title.txt")
	if err != nil {
		panic(err)
	}

	m := model{
		help:        help.InitialModel(),
		currentView: menuView,
		menuOptions: []string{
			"Local 2 Player",
			"Quit",
		},
		cursor: 0,
		title:  string(result),
	}

	if game != nil {
		m.StartGame(game)
	}

	return m
}

func (m *model) StartGame(game *mancala.Game) {
	m.board = board.InitialModel(game)
	m.currentView = boardView
}

func (m model) Init() tea.Cmd {
	helpCmd := m.help.Init()
	boardCmd := m.board.Init()

	return tea.Batch(helpCmd, boardCmd)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd, helpCmd, boardCmd tea.Cmd

	// Check for force quit
	if v, ok := msg.(tea.KeyMsg); ok && key.Matches(v, keys.Keys.ForceQuit) {
		return m, tea.Quit
	}

	switch m.currentView {
	// When in menu view
	case menuView:
		switch msg := msg.(type) {
		// When msg is a key press
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.Keys.Quit):
				return m, tea.Quit
			case key.Matches(msg, keys.Keys.Up):
				if m.cursor > 0 {
					m.cursor--
				}
			case key.Matches(msg, keys.Keys.Down):
				if m.cursor < len(m.menuOptions)-1 {
					m.cursor++
				}
			case key.Matches(msg, keys.Keys.Submit):
				switch m.cursor {
				case 0:
					p1 := mancala.NewPlayer("Player 1")
					p2 := mancala.NewPlayer("Player 2")
					game := mancala.NewGame(p1, p2, 4, 6)
					m.StartGame(game)
				case 1:
					return m, tea.Quit
				}
			}
		}
	case boardView:
		m.board, boardCmd = m.board.Update(msg)
	}

	m.help, helpCmd = m.help.Update(msg)

	return m, tea.Batch(cmd, helpCmd, boardCmd)
}

func (m model) View() string {
	if m.currentView == boardView {
		return lipgloss.JoinVertical(lipgloss.Left, m.board.View(), m.help.View())
	}

	return lipgloss.JoinVertical(lipgloss.Left, m.ViewMenu(), m.help.View())
}

func (m model) ViewMenu() string {
	var s strings.Builder

	s.WriteString(fmt.Sprintf("Welcome to\n%s\n\n", m.title))

	var optionLines []string
	for i, option := range m.menuOptions {
		cursor, style := " ", NoStyle

		if m.cursor == i {
			style = Cyan
			cursor = ">"
		}
		optionLines = append(optionLines, style(fmt.Sprintf("%s %s\n", cursor, option)))
	}
	s.WriteString(lipgloss.JoinVertical(lipgloss.Left, optionLines...))

	return s.String()
}
