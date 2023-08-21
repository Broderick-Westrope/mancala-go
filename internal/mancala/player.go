package mancala

// Player represents a player in a Mancala game.
type Player interface {
	GetName() string    // GetName returns the player's name.
	GetScore() int      // GetScore returns the player's score.
	SetScore(score int) // SetScore sets the player's score.
}

// Human represents a human player in a Mancala game. It contains a name and a score.
type Human struct {
	Name  string
	Score int
}

// GetName returns the human's name.
func (human *Human) GetName() string {
	return human.Name
}

// GetScore returns the human's score.
func (human *Human) GetScore() int {
	return human.Score
}

// SetScore sets the human's score.
func (human *Human) SetScore(score int) {
	human.Score = score
}

// NewHuman creates a new Human with the given name.
func NewHuman(name string) *Human {
	return &Human{
		Name:  name,
		Score: 0,
	}
}
