package mancala

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type BoardSide struct {
	Player *Player
	Pits   []int
	Store  int
}

func NewBoardSide(player *Player, stonesPerPit int, pitsPerSide int) *BoardSide {
	side := &BoardSide{
		Player: player,
		Store:  0,
	}

	for i := 0; i < pitsPerSide; i++ {
		side.Pits = append(side.Pits, stonesPerPit)
	}

	return side
}

func (side *BoardSide) ArePitsEmpty() bool {
	for _, pit := range side.Pits {
		if pit > 0 {
			return false
		}
	}

	return true
}

func (side *BoardSide) GetScore() int {
	var score int
	for _, pit := range side.Pits {
		score += pit
	}
	return score + side.Store
}

// Empties the pit and returns the number of stones in it
func (side *BoardSide) Capture(pitIndex int) int {
	stones := side.Pits[pitIndex]
	side.Pits[pitIndex] = 0
	return stones
}

func (side *BoardSide) ValidatePitIndex(pitIndex int) {
	if pitIndex < 0 || pitIndex >= len(side.Pits) {
		log.Error().Msg(fmt.Sprintf("Invalid pit index: %d", pitIndex))
	}
}

func (side *BoardSide) GetStones(pitIndex int) int {
	side.ValidatePitIndex(pitIndex)

	if side.Pits[pitIndex] == 0 {
		log.Error().Msgf("Pit %d is empty", pitIndex)
		return 0
	}

	stones := side.Pits[pitIndex]
	side.Pits[pitIndex] = 0
	return stones
}

// Distributes stones along board side starting at pitIndex.
//
// Returns:
//   - the number of stones that were left over after distribution
//   - whether the player gets another turn
//   - the pit index the player is capturing from (-1 for no capture)
func (side *BoardSide) DistributeStones(pitIndex int, stones int) (int, bool, int) {
	pitCount := len(side.Pits)
	for i := pitIndex; i < pitCount && stones > 0; i++ {
		// If one stone left and the next pit is empty, capture
		if stones == 1 && side.Pits[i] == 0 {
			// Add one for our stone
			side.Store++
			// Return the pit index (of the opponents board side) we want to capture from
			return 0, false, (pitCount - 1) - i
		}

		side.Pits[i]++
		stones--
	}

	switch {
	case stones == 0:
		return 0, false, -1
	case stones == 1:
		side.Store++
		return 0, true, -1
	case stones > 1:
		side.Store++
		return stones - 1, false, -1
	default:
		log.Error().Msgf("Invalid number of stones: %d", stones)
		return 0, false, -1
	}
}
