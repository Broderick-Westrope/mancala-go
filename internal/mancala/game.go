package mancala

import (
	"fmt"
)

// Game represents a game of Mancala. It contains:
// - a BoardSide for each player (Side1, Side2),
// - and an unsigned integer (Turn) representing which player's turn it is.
type Game struct {
	Side1 *BoardSide
	Side2 *BoardSide
	Turn  uint8
}

// The value of Turn should always be one of these constants.
const (
	Player1Turn = 1
	Player2Turn = 2
)

// NewGame creates a new game of Mancala with the given players, number of stones per pit, and number of pits per side.
func NewGame(player1 Player, player2 Player, stonesPerPit int, pitsPerSide int) *Game {
	return &Game{
		Side1: NewBoardSide(player1, stonesPerPit, pitsPerSide),
		Side2: NewBoardSide(player2, stonesPerPit, pitsPerSide),
		Turn:  Player1Turn,
	}
}

// ExecuteMove executes a move for the current player.
// It returns an error if the game is over or if the given pit index is invalid.
func (g *Game) ExecuteMove(pitIndex int) error {
	if g.IsOver() {
		return fmt.Errorf("game is over")
	}

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
	err := currentSide.ValidatePitIndex(pitIndex)
	if err != nil {
		return err
	}
	stones := currentSide.GetStones(pitIndex)
	if stones == 0 {
		return fmt.Errorf("pit %d is empty", pitIndex)
	}

	// Distribute stones
	pitCount := len(currentSide.Pits)
	pitIndex++
	for stones > 0 {
		for i := pitIndex; i < pitCount && stones > 0; i++ {
			currentSide.Pits[i]++
			stones--

			if currentSide == turnSide && stones == 0 && currentSide.Pits[i] == 1 {
				// Perform capture
				currentSide.Store += currentSide.GetStones(i)
				currentSide.Store += opposingSide.GetStones(currentSide.GetOpposingPitIndex(i))
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

// AlternateTurn alternates the turn between players.
func (g *Game) AlternateTurn() {
	if g.Turn == Player1Turn {
		g.Turn = Player2Turn
		return
	}
	g.Turn = Player1Turn
}

// IsOver returns true if the game is over, false otherwise.
func (g *Game) IsOver() bool {
	return g.Side1.ArePitsEmpty() || g.Side2.ArePitsEmpty()
}
