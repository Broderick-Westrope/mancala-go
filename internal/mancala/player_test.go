package mancala_test

import (
	"testing"

	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
)

func TestNewPlayer(t *testing.T) {
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
