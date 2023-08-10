package mancala

type Game struct {
	Side1 *BoardSide
	Side2 *BoardSide
	Turn  uint8
}

func (g *Game) ExecuteMove(pitIndex int) {
	ValidatePitIndex(pitIndex)

	if g.Turn == Player1Turn {
		g.Side1.ExecuteMove(pitIndex)
	} else {
		g.Side2.ExecuteMove(pitIndex)
	}
}
