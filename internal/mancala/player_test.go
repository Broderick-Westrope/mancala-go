package mancala_test

import (
	"testing"

	"github.com/Broderick-Westrope/mancala-go/internal/mancala"
	"github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
	player := mancala.NewPlayer("Test Name 1")

	assert.Equal(t, player.Name, "Test Name 1")
	assert.Equal(t, player.Score, 0)
}
