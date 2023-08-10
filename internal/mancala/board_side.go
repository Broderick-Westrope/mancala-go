package mancala

import "github.com/rs/zerolog/log"

type BoardSide struct {
	Player *Player
	Pits   []int
	Store  int
}

func NewBoardSide(player *Player) *BoardSide {
	side := &BoardSide{
		Player: player,
		Store:  0,
	}

	for i := 0; i < PitsPerSide; i++ {
		side.Pits = append(side.Pits, StonesPerPit)
	}

	return side
}

// Returns: the number of stones that were left over after the move, whether the player gets another turn, and the pit index the player is capturing from (-1 for no capture)
func (side *BoardSide) ExecuteMove(pitIndex int) (int, bool, int) {
	ValidatePitIndex(pitIndex)

	stones := side.Pits[pitIndex]

	side.Pits[pitIndex] = 0

	for i := pitIndex; i < PitsPerSide && stones > 0; i++ {
		// If one stone left and the next pit is empty, capture
		if stones == 1 && side.Pits[i] == 0 {
			// Add one for our stone
			side.Store++
			// Return the pit index (of the opponents board side) we want to capture from
			return 0, false, (PitsPerSide - 1) - i
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
