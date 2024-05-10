package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer(t *testing.T) {
	t.Run("should be able to gain experience", func(t *testing.T) {
		player := NewPlayer("", Warrior)

		assert.Equal(t, 0, player.experiencePoints)

		player.addExperiencePoints(10)

		assert.Equal(t, 10, player.experiencePoints)
	})
}
