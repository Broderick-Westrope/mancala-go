package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
	"github.com/Broderick-Westrope/mancala-go/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	name1 := flag.String("name1", "Player 1", "Name of player 1")
	name2 := flag.String("name2", "Player 2", "Name of player 2")
	mode := flag.String("mode", "local", "Game type (local, minimax)")
	pits := flag.Int("pits", 6, "Number of pits per side")
	stones := flag.Int("stones", 4, "Number of stones per pit")

	flag.Parse()

	if *pits < 1 {
		slog.Error(fmt.Sprintf("invalid number of pits: %d", *pits))
		os.Exit(1)
	}
	if *stones < 1 {
		slog.Error(fmt.Sprintf("invalid number of stones: %d", *stones))
		os.Exit(1)
	}

	var game *mancala.Game
	switch *mode {
	case "local":
		var player1, player2 *mancala.Human
		player1 = mancala.NewHuman(*name1)
		player2 = mancala.NewHuman(*name2)
		game = mancala.NewGame(player1, player2, *stones, *pits)
	case "minimax":
		var player1, player2 *mancala.MinimaxBot
		player1 = mancala.NewMinimaxBot("Minimax Bot")
		if *name2 == "Player 2" {
			player2 = mancala.NewMinimaxBot("Player")
		} else {
			player2 = mancala.NewMinimaxBot(*name2)
		}
		game = mancala.NewGame(player1, player2, *stones, *pits)
	default:
		slog.Error(fmt.Sprintf("invalid game mode: %s", *mode))
		os.Exit(1)
	}

	m := tui.InitialModel(game)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
