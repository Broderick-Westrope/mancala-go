package mancala

import (
	"fmt"
)

// BoardSide represents one side of a Mancala board. It contains a Player, a slice of integers (Pits) representing the number of stones in each pit, and an integer (Store) representing the number of stones in the store.
type BoardSide struct {
	Player Player
	Pits   []int
	Store  int
}

// NewBoardSide creates a new BoardSide with the given player, number of stones per pit, and number of pits per side.
func NewBoardSide(player Player, stonesPerPit int, pitsPerSide int) *BoardSide {
	side := &BoardSide{
		Player: player,
		Store:  0,
	}

	for i := 0; i < pitsPerSide; i++ {
		side.Pits = append(side.Pits, stonesPerPit)
	}

	return side
}

// ArePitsEmpty returns true if all pits are empty.
func (side *BoardSide) ArePitsEmpty() bool {
	for _, pit := range side.Pits {
		if pit > 0 {
			return false
		}
	}

	return true
}

// GetScore returns the total number of stones on this side (pits and store).
func (side *BoardSide) GetScore() int {
	var score int
	for _, pit := range side.Pits {
		score += pit
	}
	return score + side.Store
}

// ValidatePitIndex returns an error if the given pit index is invalid. Otherwise, it returns nil.
func (side *BoardSide) ValidatePitIndex(pitIndex int) error {
	if pitIndex < 0 || pitIndex >= len(side.Pits) {
		return fmt.Errorf("invalid pit index: %d", pitIndex)
	}
	return nil
}

// GetStones returns the number of stones in the given pit and removes them from the pit.
func (side *BoardSide) GetStones(pitIndex int) int {
	stones := side.Pits[pitIndex]
	side.Pits[pitIndex] = 0
	return stones
}

// GetOpposingPitIndex returns the index of the pit on the opposing side of the board. This takes into account the pits being reversed for the opposing player.
func (side *BoardSide) GetOpposingPitIndex(pitIndex int) int {
	return (len(side.Pits) - 1) - pitIndex
}
