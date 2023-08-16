package main

import (
	"flag"

	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
	"github.com/Broderick-Westrope/mancala-go/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	name1 := flag.String("name1", "Player 2", "Name of player 1")
	name2 := flag.String("name2", "Player 1", "Name of player 2")
	mode := flag.String("mode", "local", "Game type (local)")
	pits := flag.Int("pits", 6, "Number of pits per side")
	stones := flag.Int("stones", 4, "Number of stones per pit")

	flag.Parse()

	var game *mancala.Game
	switch *mode {
	case "local":
		var player1, player2 *mancala.Player
		if *name1 != "" {
			player1 = mancala.NewPlayer(*name1)
		} else {
			player1 = mancala.NewPlayer("Player 1")
		}
		if *name2 != "" {
			player2 = mancala.NewPlayer(*name2)
		} else {
			player2 = mancala.NewPlayer("Player 2")
		}
		game = mancala.NewGame(player1, player2, *stones, *pits)
	}

	m := tui.InitialModel(game)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
