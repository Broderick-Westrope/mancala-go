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
			if ans.Name != tt.expectedName {
				t.Errorf("got %s, want %s", ans.Name, tt.expectedName)
			}
			if ans.Score != tt.expectedScore {
				t.Errorf("got %d, want %d", ans.Score, tt.expectedScore)
			}
		})
	}
}
