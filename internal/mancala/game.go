package mancala

import (
	"fmt"
)

type Game struct {
	Side1 *BoardSide
	Side2 *BoardSide
	Turn  uint8
}

const (
	Player1Turn = 1
	Player2Turn = 2
)

func NewGame(player1 *Player, player2 *Player, stonesPerPit int, pitsPerSide int) *Game {
	return &Game{
		Side1: NewBoardSide(player1, stonesPerPit, pitsPerSide),
		Side2: NewBoardSide(player2, stonesPerPit, pitsPerSide),
		Turn:  Player1Turn,
	}
}

func (g *Game) ExecuteMove(pitIndex int) error {
	// Get correct board side
	var currentSide, opposingSide *BoardSide
	if g.Turn == Player1Turn {
		currentSide = g.Side1
		opposingSide = g.Side2
	} else {
		currentSide = g.Side2
		opposingSide = g.Side1
	}
	turnSide := currentSide

	// Get stones from pit
	currentSide.ValidatePitIndex(pitIndex)
	stones := currentSide.GetStones(pitIndex)
	if stones == 0 {
		return fmt.Errorf("pit %d is empty", pitIndex)
	}

	// Distribute stones
	pitCount := len(currentSide.Pits)
	pitIndex++
	for stones > 0 {
		// TODO: try to use pitIndex++ instead of pitIndex = pitIndex + 1
		for i := pitIndex; i < pitCount && stones > 0; i++ {
			currentSide.Pits[i]++
			stones--

			if currentSide == turnSide && stones == 0 && currentSide.Pits[i] == 1 {
				// Perform capture
				currentSide.Store += currentSide.Capture(i)
				currentSide.Store += opposingSide.Capture(currentSide.GetOpposingPitIndex(i))
				g.AlternateTurn()
				return nil
			}
		}

		if stones > 0 && currentSide == turnSide {
			currentSide.Store++
			stones--
			// If last stone lands in players store, they get another turn
			if stones == 0 {
				return nil
			}
		}

		if stones < 0 {
			return fmt.Errorf("stones is negative: %d", stones)
		}

		// Switch sides and continue distributing
		currentSide, opposingSide = opposingSide, currentSide
		pitIndex = 0
	}

	g.AlternateTurn()
	return nil
}

func (g *Game) AlternateTurn() {
	if g.Turn == Player1Turn {
		g.Turn = Player2Turn
	} else {
		g.Turn = Player1Turn
	}
}
