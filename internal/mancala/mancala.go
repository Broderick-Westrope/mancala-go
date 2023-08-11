package mancala

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
