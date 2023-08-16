package mancala

type Player interface {
	GetName() string
	GetScore() int
	SetScore(score int)
}

type Human struct {
	Name  string
	Score int
}

func (human *Human) GetName() string {
	return human.Name
}

func (human *Human) GetScore() int {
	return human.Score
}

func (human *Human) SetScore(score int) {
	human.Score = score
}

func NewHuman(name string) *Human {
	return &Human{
		Name:  name,
		Score: 0,
	}
}
