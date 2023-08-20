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

	index, value := -1, math.MinInt

	for i, pit := range sim.Side1.Pits {
		if pit > 0 {
			newSim := sim.copy()
			newSim.ExecuteMove(i)
			newValue := newSim.getMove(15, game.Turn, math.MinInt, math.MaxInt)
			if newValue > value {
				index = i
				value = newValue
			}
		}
	}

	return index
}

func (sim *Game) getMove(depth int, maximiser uint8, alpha, beta int) int {
	if sim.IsOver() || depth == 0 {
		return sim.Side1.Store - sim.Side2.Store
	}
	if sim.Turn == maximiser {
		bestValue := math.MinInt
		for i, pit := range sim.Side1.Pits {
			if pit > 0 {
				newSim := sim.copy()
				newSim.ExecuteMove(i)
				value := newSim.getMove(depth-1, maximiser, alpha, beta)
				bestValue = max(bestValue, value)
				alpha = max(alpha, bestValue)
				if beta <= alpha {
					break
				}
			}
		}
		return bestValue
	} else {
		bestValue := math.MaxInt
		for i, pit := range sim.Side2.Pits {
			if pit > 0 {
				newSim := sim.copy()
				newSim.ExecuteMove(i)
				value := newSim.getMove(depth-1, maximiser, alpha, beta)
				bestValue = min(bestValue, value)
				beta = min(beta, bestValue)
				if beta <= alpha {
					break
				}
			}
		}
		return bestValue
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
