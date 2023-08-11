package mancala_test

import (
	"testing"

	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
)

func TestGame_ExecuteMove(t *testing.T) {
	tests := []struct {
		name          string
		game          *mancala.Game
		pitIndex      int
		expectedGame  *mancala.Game
		expectedAnErr bool
	}{
		{
			name: "Chosen pit is empty",
			game: &mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{4, 0, 4, 4, 4, 4},
					Store: 1,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 4, 4},
					Store: 1,
				},
				Turn: mancala.Player1Turn,
			},
			pitIndex: 1,
			expectedGame: &mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{4, 0, 4, 4, 4, 4},
					Store: 1,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 4, 4},
					Store: 1,
				},
				Turn: mancala.Player1Turn,
			},
			expectedAnErr: true,
		},
		{
			name: "Move ending on players pit",
			game: &mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 4, 4},
					Store: 2,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 4, 4},
					Store: 2,
				},
				Turn: mancala.Player1Turn,
			},
			pitIndex: 0,
			expectedGame: &mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{0, 5, 5, 5, 5, 4},
					Store: 2,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 4, 4},
					Store: 2,
				},
				Turn: mancala.Player2Turn,
			},
		},
		{
			name: "Move ending on players store",
			game: &mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 4, 4},
					Store: 3,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 4, 4},
					Store: 3,
				},
				Turn: mancala.Player1Turn,
			},
			pitIndex: 2,
			expectedGame: &mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{4, 4, 0, 5, 5, 5},
					Store: 4,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 4, 4},
					Store: 3,
				},
				Turn: mancala.Player1Turn,
			},
		},
		{
			name: "Move ending on opponents full pit",
			game: &mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 4, 4},
					Store: 3,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 4, 4},
					Store: 3,
				},
				Turn: mancala.Player1Turn,
			},
			pitIndex: 5,
			expectedGame: &mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 4, 0},
					Store: 4,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{5, 5, 5, 4, 4, 4},
					Store: 3,
				},
				Turn: mancala.Player2Turn,
			},
		},
		{
			name: "Move ending on opponents empty pit",
			game: &mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 4, 4},
					Store: 3,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{0, 0, 0, 0, 0, 0},
					Store: 3,
				},
				Turn: mancala.Player1Turn,
			},
			pitIndex: 5,
			expectedGame: &mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 4, 0},
					Store: 4,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{1, 1, 1, 0, 0, 0},
					Store: 3,
				},
				Turn: mancala.Player2Turn,
			},
		},
		{
			name: "Move ending on opponents store",
			game: &mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 4, 8},
					Store: 3,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 4, 4},
					Store: 3,
				},
				Turn: mancala.Player1Turn,
			},
			pitIndex: 5,
			expectedGame: &mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{5, 4, 4, 4, 4, 0},
					Store: 4,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{5, 5, 5, 5, 5, 5},
					Store: 3,
				},
				Turn: mancala.Player2Turn,
			},
		},
		{
			name: "Move ending on players empty pit",
			game: &mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 0, 4},
					Store: 3,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{4, 4, 4, 4, 4, 4},
					Store: 3,
				},
				Turn: mancala.Player1Turn,
			},
			pitIndex: 0,
			expectedGame: &mancala.Game{
				Side1: &mancala.BoardSide{
					Pits:  []int{0, 5, 5, 5, 0, 4},
					Store: 8,
				},
				Side2: &mancala.BoardSide{
					Pits:  []int{4, 0, 4, 4, 4, 4},
					Store: 3,
				},
				Turn: mancala.Player2Turn,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.game.ExecuteMove(tt.pitIndex)
			if err != nil && !tt.expectedAnErr {
				t.Errorf("Unexpected error: %v", err)
			} else if err == nil && tt.expectedAnErr {
				t.Errorf("Expected error but got none")
			}
			checkEquals(t, "Turn", tt.game.Turn, tt.expectedGame.Turn)
			checkEquals(t, "Player 1 Pits", tt.game.Side1.Pits, tt.expectedGame.Side1.Pits)
			checkEquals(t, "Player 1 Store", tt.game.Side1.Store, tt.expectedGame.Side1.Store)
			checkEquals(t, "Player 2 Pits", tt.game.Side2.Pits, tt.expectedGame.Side2.Pits)
			checkEquals(t, "Player 2 Store", tt.game.Side2.Store, tt.expectedGame.Side2.Store)
		})
	}
}
