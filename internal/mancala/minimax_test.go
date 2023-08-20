package mancala_test

import (
	"testing"

	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
)

func TestNewMinimaxBot(t *testing.T) {
	tests := []struct {
		name          string
		playerName    string
		expectedName  string
		expectedScore int
	}{
		{"Valid name, valid score", "Test Name 1", "Test Name 1", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := mancala.NewMinimaxBot(tt.playerName)
			checkEquals(t, "Name", ans.Name, tt.expectedName)
			checkEquals(t, "Score", ans.Score, tt.expectedScore)
		})
	}
}

func TestMinimaxBot_GetName_GetScore(t *testing.T) {
	tests := []struct {
		name          string
		bot           *mancala.MinimaxBot
		expectedName  string
		expectedScore int
	}{
		{
			"Valid name",
			&mancala.MinimaxBot{
				Name:  "Test Name 1",
				Score: 5,
			},
			"Test Name 1",
			5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := tt.bot.GetName()
			score := tt.bot.GetScore()
			checkEquals(t, "Name", name, tt.expectedName)
			checkEquals(t, "Score", score, tt.expectedScore)
		})
	}
}

func TestMinimaxBot_SetScore(t *testing.T) {
	tests := []struct {
		name          string
		bot           *mancala.MinimaxBot
		expectedScore int
	}{
		{
			"Valid score",
			&mancala.MinimaxBot{
				Name:  "Test Name 1",
				Score: 5,
			},
			12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bot.SetScore(tt.expectedScore)
			checkEquals(t, "Score", tt.bot.Score, tt.expectedScore)
		})
	}
}

func TestMinimaxBot_GetMove(t *testing.T) {
	tests := []struct {
		name         string
		game         *mancala.Game
		expectedMove int
	}{
		{
			"Valid move 1",
			&mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{0, 0, 4, 0, 0, 0},
					Store: 2,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{0, 5, 0, 0, 0, 0},
					Store: 2,
				},
				Turn: mancala.Player1Turn,
			},
			2,
		},
		{
			"Valid move 2",
			&mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{0, 8, 12, 1, 0, 4},
					Store: 2,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{0, 5, 0, 3, 4, 0},
					Store: 2,
				},
				Turn: mancala.Player2Turn,
			},
			1,
		},
		{
			"Steal 1",
			&mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{0, 0, 0, 1, 0, 0},
					Store: 2,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{0, 5, 0, 3, 4, 0},
					Store: 2,
				},
				Turn: mancala.Player1Turn,
			},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot := mancala.NewMinimaxBot("Test Bot")
			move := bot.GetMove(tt.game)
			checkEquals(t, "Move", move, tt.expectedMove)
		})
	}
}
