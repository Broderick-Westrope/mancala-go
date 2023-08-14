package menu

import (
	"fmt"
	"os"
	"strings"

	"github.com/Broderick-Westrope/mancala-go/internal/tui/keys"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	options []string
	cursor  int
	title   string
}

func InitialModel() Model {
	result, err := os.ReadFile("assets/mancala_title.txt")
	if err != nil {
		panic(err)
	}

	return Model{
		options: []string{
			"Local 2 Player",
			"Quit",
		},
		cursor: 0,
		title:  string(result),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Up):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, keys.Keys.Down):
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case key.Matches(msg, keys.Keys.Submit):
			switch m.cursor {
			case 1:
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	var s strings.Builder

	s.WriteString(fmt.Sprintf("Welcome to\n%s\n\n", m.title))

	var pararaphs []string
	for i, option := range m.options {
		cursor, style := " ", NoStyle

		if m.cursor == i {
			style = Cyan
			cursor = ">"
		}
		pararaphs = append(pararaphs, style(fmt.Sprintf("%s %s\n", cursor, option)))
	}
	s.WriteString(lipgloss.JoinVertical(lipgloss.Left, pararaphs...))

	return s.String()
}
