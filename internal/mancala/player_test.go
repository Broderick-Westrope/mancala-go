package mancala_test

import (
	"testing"

	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
)

func TestNewHuman(t *testing.T) {
	tests := []struct {
		name          string
		playerName    string
		expectedName  string
		expectedScore int
	}{
		{"Valid name, valid score", "Test Name 1", "Test Name 1", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := mancala.NewHuman(tt.playerName)
			checkEquals(t, "Name", ans.Name, tt.expectedName)
			checkEquals(t, "Score", ans.Score, tt.expectedScore)
		})
	}
}

func TestHuman_GetName_GetScore(t *testing.T) {
	tests := []struct {
		name          string
		human         *mancala.Human
		expectedName  string
		expectedScore int
	}{
		{
			"Valid name",
			&mancala.Human{
				Name:  "Test Name 1",
				Score: 5,
			},
			"Test Name 1",
			5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := tt.human.GetName()
			score := tt.human.GetScore()
			checkEquals(t, "Name", name, tt.expectedName)
			checkEquals(t, "Score", score, tt.expectedScore)
		})
	}
}

func TestHuman_SetScore(t *testing.T) {
	tests := []struct {
		name          string
		human         *mancala.Human
		expectedScore int
	}{
		{
			"Valid score",
			&mancala.Human{
				Name:  "Test Name 1",
				Score: 5,
			},
			12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.human.SetScore(tt.expectedScore)
			checkEquals(t, "Score", tt.human.Score, tt.expectedScore)
		})
	}
}
