package mancala_test

import (
	"reflect"
	"testing"

	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
)

func TestNewBoardSide(t *testing.T) {
	tests := []struct {
		name           string
		player         *mancala.Player
		stonesPerPit   int
		pitsPerSide    int
		expectedPlayer *mancala.Player
		expectedPits   []int
		expectedStore  int
	}{
		{
			"Valid player, 4 stones, 6 pits",
			mancala.NewPlayer("Test Name 1"),
			4, 6,
			mancala.NewPlayer("Test Name 1"),
			[]int{4, 4, 4, 4, 4, 4},
			0,
		},
		{
			"Valid player, 2 stones, 3 pits",
			mancala.NewPlayer("Test Name 1"),
			2, 3,
			mancala.NewPlayer("Test Name 1"),
			[]int{2, 2, 2},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := mancala.NewBoardSide(tt.player, tt.stonesPerPit, tt.pitsPerSide)
			if *ans.Player != *tt.expectedPlayer {
				t.Errorf("got %v, want %v", ans.Player, tt.expectedPlayer)
			}
			if ans.Store != tt.expectedStore {
				t.Errorf("got %v, want %v", ans.Store, tt.expectedStore)
			}
			if !reflect.DeepEqual(ans.Pits, tt.expectedPits) {
				t.Errorf("got %v, want %v", ans.Pits, tt.expectedPits)
			}
		})
	}
}

func TestArePitsEmpty(t *testing.T) {
	tests := []struct {
		name          string
		player        *mancala.Player
		stonesPerPit  int
		pitsPerSide   int
		expectedEmpty bool
	}{
		{
			"4 stones, 6 pits, not empty",
			mancala.NewPlayer("Test Name 1"),
			4, 6, false,
		},
		{
			"0 stones, 6 pits, empty",
			mancala.NewPlayer("Test Name 1"),
			0, 6, true,
		},
		{
			"6 stones, 0 pits, empty",
			mancala.NewPlayer("Test Name 1"),
			6, 0, true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			side := mancala.NewBoardSide(tt.player, tt.stonesPerPit, tt.pitsPerSide)
			ans := side.ArePitsEmpty()
			if ans != tt.expectedEmpty {
				t.Errorf("got %v, want %v", ans, tt.expectedEmpty)
			}
		})
	}
}

func TestGetScore(t *testing.T) {
	tests := []struct {
		name          string
		player        *mancala.Player
		pits          []int
		store         int
		expectedScore int
	}{
		{
			"5 pits, 1 stone each, 6 in store",
			mancala.NewPlayer("Test Name 1"),
			[]int{1, 1, 1, 1, 1},
			6,
			11,
		},
		{
			"8 pits, incremental stones, 0 in store",
			mancala.NewPlayer("Test Name 1"),
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			0,
			36,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			side := mancala.NewBoardSide(tt.player, 4, 6)
			side.Pits = tt.pits
			side.Store = tt.store
			ans := side.GetScore()
			if ans != tt.expectedScore {
				t.Errorf("got %v, want %v", ans, tt.expectedScore)
			}
		})
	}
}

func TestCapture(t *testing.T) {
	tests := []struct {
		name                   string
		player                 *mancala.Player
		pits                   []int
		captureIndex           int
		expectedCapturedStones int
	}{
		{
			"Capture 5 stones, index 3",
			mancala.NewPlayer("Test Name 1"),
			[]int{1, 0, 7, 5, 9, 12},
			3, 5,
		},
		{
			"Capture 9 stones, index 4",
			mancala.NewPlayer("Test Name 1"),
			[]int{1, 0, 7, 5, 9, 12},
			4, 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			side := mancala.NewBoardSide(tt.player, 4, 6)
			side.Pits = tt.pits
			ans := side.Capture(tt.captureIndex)
			if ans != tt.expectedCapturedStones {
				t.Errorf("got %v, want %v", ans, tt.expectedCapturedStones)
			}
			if side.Pits[tt.captureIndex] != 0 {
				t.Errorf("got %v, want %d", side.Pits[tt.captureIndex], 0)
			}
		})
	}
}
