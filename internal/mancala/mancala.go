package mancala

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

const (
	Player1Turn = 1
	Player2Turn = 2
)

var (
	StonesPerPit = 4
	PitsPerSide  = 6
)

func NewGame(player1 *Player, player2 *Player) *Game {
	return &Game{
		Side1: NewBoardSide(player1),
		Side2: NewBoardSide(player2),
		Turn:  Player1Turn,
	}
}

func ValidatePitIndex(pitIndex int) {
	if pitIndex < 0 || pitIndex >= PitsPerSide {
		log.Error().Msg(fmt.Sprintf("Invalid pit index: %d", pitIndex))
	}
}
