package main

import (
	"github.com/Broderick-Westrope/mancala-go/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := tui.InitialModel(nil)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
