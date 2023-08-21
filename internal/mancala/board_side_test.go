package mancala_test

import (
	"testing"

	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
)

func TestNewBoardSide(t *testing.T) {
	tests := []struct {
		name           string
		player         mancala.Player
		stonesPerPit   int
		pitsPerSide    int
		expectedPlayer mancala.Player
		expectedPits   []int
		expectedStore  int
	}{
		{
			"Valid player, 4 stones, 6 pits",
			mancala.NewHuman("Test Name 1"),
			4, 6,
			mancala.NewHuman("Test Name 1"),
			[]int{4, 4, 4, 4, 4, 4},
			0,
		},
		{
			"Valid player, 2 stones, 3 pits",
			mancala.NewHuman("Test Name 1"),
			2, 3,
			mancala.NewHuman("Test Name 1"),
			[]int{2, 2, 2},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := mancala.NewBoardSide(tt.player, tt.stonesPerPit, tt.pitsPerSide)
			checkEquals(t, "Player", ans.Player, tt.expectedPlayer)
			checkEquals(t, "Pits", ans.Pits, tt.expectedPits)
			checkEquals(t, "Store", ans.Store, tt.expectedStore)
		})
	}
}

func TestBoardSide_GetStones(t *testing.T) {
	tests := []struct {
		name               string
		side               *mancala.BoardSide
		pitIndex           int
		expectedStonesLeft int
	}{
		{
			"Chosen pit is empty",
			&mancala.BoardSide{
				Player: mancala.NewHuman("Test Name 1"),
				Pits:   []int{0, 4, 4, 4, 4, 4},
				Store:  1,
			},
			0, 0,
		},
		{
			"Chosen pit has stones",
			&mancala.BoardSide{
				Player: mancala.NewHuman("Test Name 1"),
				Pits:   []int{4, 4, 4, 4, 4, 4},
				Store:  1,
			},
			5, 4,
		},
		{
			"Capture 5 stones, index 3",
			&mancala.BoardSide{Pits: []int{1, 0, 7, 5, 9, 12}},
			3, 5,
		},
		{
			"Capture 9 stones, index 4",
			&mancala.BoardSide{Pits: []int{1, 0, 7, 5, 9, 12}},
			4, 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.side.GetStones(tt.pitIndex)
			checkEquals(t, "Stones Returned", ans, tt.expectedStonesLeft)
			checkEquals(t, "Stones At Index", tt.side.Pits[tt.pitIndex], 0)
		})
	}
}

func TestBoardSide_ArePitsEmpty(t *testing.T) {
	tests := []struct {
		name          string
		side          *mancala.BoardSide
		expectedEmpty bool
	}{
		{
			"4 stones, 6 pits, not empty",
			&mancala.BoardSide{
				Pits: []int{4, 4, 4, 4, 4, 4},
			},
			false,
		},
		{
			"0 stones, 6 pits, empty",
			&mancala.BoardSide{
				Pits: []int{0, 0, 0, 0, 0, 0},
			},
			true,
		},
		{
			"0 pits, empty",
			&mancala.BoardSide{
				Pits: []int{},
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.side.ArePitsEmpty()
			checkEquals(t, "Are Pits Empty", ans, tt.expectedEmpty)
		})
	}
}

func TestBoardSide_GetScore(t *testing.T) {
	tests := []struct {
		name          string
		side          *mancala.BoardSide
		expectedScore int
	}{
		{
			"5 pits, 1 stone each, 6 in store",
			&mancala.BoardSide{
				Pits:  []int{1, 1, 1, 1, 1},
				Store: 6,
			},
			11,
		},
		{
			"8 pits, incremental stones, 0 in store",
			&mancala.BoardSide{
				Pits:  []int{1, 2, 3, 4, 5, 6, 7, 8},
				Store: 0,
			},

			36,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.side.GetScore()
			checkEquals(t, "Score", ans, tt.expectedScore)
		})
	}
}

func TestBoardSide_GetOpposingPitIndex(t *testing.T) {
	tests := []struct {
		name          string
		side          *mancala.BoardSide
		pitIndex      int
		expectedIndex int
	}{
		{
			"Index 0, 6 pits",
			&mancala.BoardSide{
				Pits: []int{1, 0, 7, 5, 9, 12},
			},
			0, 5,
		},
		{
			"Index 1, 12 pits",
			&mancala.BoardSide{
				Pits: []int{1, 0, 7, 5, 9, 12, 1, 0, 7, 5, 9, 12},
			},
			1, 10,
		},
		{
			"Index 2, 4 pits",
			&mancala.BoardSide{
				Pits: []int{1, 0, 7, 5},
			},
			2, 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkEquals(t, "Opposing Pit Index", tt.side.GetOpposingPitIndex(tt.pitIndex), tt.expectedIndex)
		})
	}
}
