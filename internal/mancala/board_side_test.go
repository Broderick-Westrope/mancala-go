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

func TestBoardSide_ExecuteMove(t *testing.T) {
	tests := []struct {
		name                 string
		side                 *mancala.BoardSide
		pitIndex             int
		expectedPits         []int
		expectedStore        int
		expectedStonesLeft   int
		expectedAnotherTurn  bool
		expectedCaptureIndex int
	}{
		{
			"Chosen pit is empty",
			&mancala.BoardSide{
				Player: mancala.NewPlayer("Test Name 1"),
				Pits:   []int{0, 4, 4, 4, 4, 4},
				Store:  1,
			},
			0,
			[]int{0, 4, 4, 4, 4, 4},
			1, 0, true, -1,
		},
		{
			"Move ending on players pit",
			&mancala.BoardSide{
				Player: mancala.NewPlayer("Test Name 1"),
				Pits:   []int{4, 4, 4, 4, 4, 4},
				Store:  1,
			},
			1,
			[]int{4, 0, 5, 5, 5, 5},
			1, 0, false, -1,
		},
		{
			"Move ending on players store",
			&mancala.BoardSide{
				Player: mancala.NewPlayer("Test Name 1"),
				Pits:   []int{4, 4, 4, 4, 4, 4},
				Store:  1,
			},
			2,
			[]int{4, 4, 0, 5, 5, 5},
			2, 0, true, -1,
		},
		{
			"Move ending on opponents side",
			&mancala.BoardSide{
				Player: mancala.NewPlayer("Test Name 1"),
				Pits:   []int{4, 4, 4, 4, 4, 4},
				Store:  1,
			},
			4,
			[]int{4, 4, 4, 4, 0, 5},
			2, 2, false, -1,
		},
		{
			"Perform a capture",
			&mancala.BoardSide{
				Player: mancala.NewPlayer("Test Name 1"),
				Pits:   []int{4, 3, 4, 4, 0, 4},
				Store:  1,
			},
			1,
			[]int{4, 0, 5, 5, 0, 4},
			2, 0, false, 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ansStonesLeft, ansAnotherTurn, ansCaptureIndex := tt.side.ExecuteMove(tt.pitIndex)
			if !reflect.DeepEqual(tt.side.Pits, tt.expectedPits) {
				t.Errorf("Pits: got %v, want %v", tt.side.Pits, tt.expectedPits)
			}
			if tt.side.Store != tt.expectedStore {
				t.Errorf("Store: got %v, want %v", tt.side.Store, tt.expectedStore)
			}
			if ansStonesLeft != tt.expectedStonesLeft {
				t.Errorf("Stones Left: got %v, want %v", ansStonesLeft, tt.expectedStonesLeft)
			}
			if ansAnotherTurn != tt.expectedAnotherTurn {
				t.Errorf("Another Turn: got %v, want %v", ansAnotherTurn, tt.expectedAnotherTurn)
			}
			if ansCaptureIndex != tt.expectedCaptureIndex {
				t.Errorf("Capture Index: got %v, want %v", ansCaptureIndex, tt.expectedCaptureIndex)
			}
		})
	}
}

func TestBoardSide_ArePitsEmpty(t *testing.T) {
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

func TestBoardSide_GetScore(t *testing.T) {
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

func TestBoardSide_Capture(t *testing.T) {
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
