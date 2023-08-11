package mancala

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

func (g *Game) ExecuteMove(pitIndex int) {
	// Get correct board side
	var currentSide *BoardSide
	if g.Turn == Player1Turn {
		currentSide = g.Side1
	} else {
		currentSide = g.Side2
	}

	currentSide.ValidatePitIndex(pitIndex)

	var stones, captureIndex int
	var extraTurn bool
	// Do while there are stones left to distribute
	for i := true; i; i = stones > 0 {
		stones, extraTurn, captureIndex = currentSide.ExecuteMove(pitIndex)

		// If we have finished distributing stones
		if stones <= 0 {
			// If they get to capture
			if captureIndex >= 0 {
				// Get the other board side
				var otherSide *BoardSide
				if g.Turn == Player1Turn {
					otherSide = g.Side2
				} else {
					otherSide = g.Side1
				}
				// Add captured stones to store
				currentSide.Store += otherSide.Capture(captureIndex)
			}
			// If they don't get to take another turn, alternate the turn
			if !extraTurn {
				g.AlternateTurn()
			}
			break
		}

		// Else, we need to distribute the remaining stones on the other side
		if g.Turn == Player1Turn {
			currentSide = g.Side2
		} else {
			currentSide = g.Side1
		}
	}
}

func (g *Game) AlternateTurn() {
	if g.Turn == Player1Turn {
		g.Turn = Player2Turn
	} else {
		g.Turn = Player1Turn
	}
}
