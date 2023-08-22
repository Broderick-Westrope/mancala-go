package mancala

import (
	"fmt"
	"log/slog"
	"math"
)

const (
	maxDepth = 15
)

// MinimaxBot represents a bot that uses the minimax algorithm to determine its next move.
// It contains a name and a score.
type MinimaxBot struct {
	Name  string
	Score int
}

// GetName returns the bot's name.
func (bot *MinimaxBot) GetName() string {
	return bot.Name
}

// GetScore returns the bot's score.
func (bot *MinimaxBot) GetScore() int {
	return bot.Score
}

// SetScore sets the bot's score.
func (bot *MinimaxBot) SetScore(score int) {
	bot.Score = score
}

// NewMinimaxBot creates a new MinimaxBot with the given name.
func NewMinimaxBot(name string) *MinimaxBot {
	return &MinimaxBot{
		Name:  name,
		Score: 0,
	}
}

// GetMove returns the bot's next move using the minimax algorithm.
func (bot *MinimaxBot) GetMove(game *Game) int {
	var maxSide, minSide *BoardSide
	if game.Turn == Player1Turn {
		maxSide = game.Side1
		minSide = game.Side2
	} else {
		maxSide = game.Side2
		minSide = game.Side1
	}

	g := &Game{
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

	for i, pit := range g.Side1.Pits {
		if pit > 0 {
			newSim := g.copy()
			err := newSim.ExecuteMove(i)
			if err != nil {
				slog.Error(fmt.Sprintf("Error executing move: %s", err.Error()))
			}
			newValue := newSim.getMove(maxDepth, game.Turn, math.MinInt, math.MaxInt)
			if newValue > value {
				index = i
				value = newValue
			}
		}
	}

	return index
}

// getMove returns the best move for the given side using the minimax algorithm.
func (g *Game) getMove(depth int, maximiser uint8, alpha, beta int) int {
	if g.IsOver() || depth == 0 {
		return g.Side1.Store - g.Side2.Store
	}
	if g.Turn == maximiser {
		return g.maximiserMove(depth, maximiser, alpha, beta)
	}
	return g.minimiserMove(depth, maximiser, alpha, beta)
}

// maximiserMove returns the best move for the maximiser.
func (g *Game) maximiserMove(depth int, maximiser uint8, alpha, beta int) int {
	bestValue := math.MinInt

	for i, pit := range g.Side1.Pits {
		if pit > 0 {
			newSim := g.copy()
			err := newSim.ExecuteMove(i)
			if err != nil {
				slog.Error(fmt.Sprintf("Error executing move: %s", err.Error()))
			}
			value := newSim.getMove(depth-1, maximiser, alpha, beta)
			bestValue = max(bestValue, value)
			alpha = max(alpha, bestValue)
			if beta <= alpha {
				break
			}
		}
	}
	return bestValue
}

// minimiserMove returns the best move for the minimiser.
func (g *Game) minimiserMove(depth int, maximiser uint8, alpha, beta int) int {
	bestValue := math.MaxInt
	for i, pit := range g.Side2.Pits {
		if pit > 0 {
			newSim := g.copy()
			err := newSim.ExecuteMove(i)
			if err != nil {
				slog.Error(fmt.Sprintf("Error executing move: %s", err.Error()))
			}
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

// copy returns a copy of the Game.
func (g *Game) copy() *Game {
	newSim := &Game{
		Side1: &BoardSide{
			Store: g.Side1.Store,
			Pits:  make([]int, len(g.Side1.Pits)),
		},
		Side2: &BoardSide{
			Store: g.Side1.Store,
			Pits:  make([]int, len(g.Side2.Pits)),
		},
		Turn: g.Turn,
	}

	copy(newSim.Side1.Pits, g.Side1.Pits)
	copy(newSim.Side2.Pits, g.Side2.Pits)

	return newSim
}
