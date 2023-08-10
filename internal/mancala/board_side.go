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

// Returns the number of stones that were left over after the move and a boolean to represent whether the player gets another turn.
func (side *BoardSide) ExecuteMove(pitIndex int) (int, bool) {
	ValidatePitIndex(pitIndex)

	stones := side.Pits[pitIndex]

	side.Pits[pitIndex] = 0

	for i := pitIndex; i < PitsPerSide && stones > 0; i++ {
		side.Pits[i]++
		stones--
	}

	switch {
	case stones == 0:
		return 0, false
	case stones == 1:
		side.Store++
		return 0, true
	case stones > 1:
		side.Store++
		return stones - 1, false
	default:
		log.Error().Msgf("Invalid number of stones: %d", stones)
		return 0, false
	}
}
