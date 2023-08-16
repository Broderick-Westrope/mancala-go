package mancala

import (
	"math"
)

type MinimaxBot struct {
	Name  string
	Score int
}

func (bot *MinimaxBot) GetName() string {
	return bot.Name
}

func (bot *MinimaxBot) GetScore() int {
	return bot.Score
}

func (bot *MinimaxBot) SetScore(score int) {
	bot.Score = score
}

func NewMinimaxBot(name string) *MinimaxBot {
	return &MinimaxBot{
		Name:  name,
		Score: 0,
	}
}

func (bot *MinimaxBot) GetMove(game *Game) int {
	var maxSide, minSide *BoardSide
	if game.Turn == Player1Turn {
		maxSide = game.Side1
		minSide = game.Side2
	} else {
		maxSide = game.Side2
		minSide = game.Side1
	}

	sim := &Game{
		Side1: &BoardSide{
			Store: maxSide.Store,
			Pits:  maxSide.Pits,
		},
		Side2: &BoardSide{
			Store: minSide.Store,
			Pits:  minSide.Pits,
		},
		Turn: game.Turn,
	}

	var index, value int

	for i, pit := range sim.Side1.Pits {
		if pit > 0 {
			newSim := sim.copy()
			newSim.ExecuteMove(i)
			newValue := newSim.getMove(5, game.Turn)
			if newValue > value {
				index = i
				value = newValue
			}
		}
	}

	return index
}

func (sim *Game) getMove(depth int, maximiser uint8) int {
	if sim.IsOver() || depth == 0 {
		return sim.Side1.Store - sim.Side2.Store
	}
	if sim.Turn == maximiser {
		value := math.MinInt
		for i, pit := range sim.Side1.Pits {
			if pit > 0 {
				newSim := sim.copy()
				newSim.ExecuteMove(i)
				value = max(value, newSim.getMove(depth-1, maximiser))
			}
		}
		return value
	} else {
		value := math.MaxInt
		for i, pit := range sim.Side2.Pits {
			if pit > 0 {
				newSim := sim.copy()
				newSim.ExecuteMove(i)
				value = max(value, newSim.getMove(depth-1, maximiser))
			}
		}
		return value
	}
}

func (sim *Game) copy() *Game {
	newSim := &Game{
		Side1: &BoardSide{
			Store: sim.Side1.Store,
			Pits:  make([]int, len(sim.Side1.Pits)),
		},
		Side2: &BoardSide{
			Store: sim.Side1.Store,
			Pits:  make([]int, len(sim.Side2.Pits)),
		},
		Turn: sim.Turn,
	}

	copy(newSim.Side1.Pits, sim.Side1.Pits)
	copy(newSim.Side2.Pits, sim.Side2.Pits)

	return newSim
}
