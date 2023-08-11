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
	stones := side.Pits[pitIndex]
	side.Pits[pitIndex] = 0
	return stones
}

func (side *BoardSide) GetOpposingPitIndex(pitIndex int) int {
	return (len(side.Pits) - 1) - pitIndex
}
