package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombat(t *testing.T) {
	t.Run("should be able to enter combat with a single enemy", func(t *testing.T) {
		player := NewPlayer("", Warrior, WithClassBasedStartingInventory())
		enemy := NewEnemy(Goblin, WithLevel(1))

		preCombatPlayerHP := player.healthPoints
		preCombatPlayerXP := player.experiencePoints
		Combat(player, enemy)

		assert.Less(t, player.healthPoints, preCombatPlayerHP)
		assert.Greater(t, player.experiencePoints, preCombatPlayerXP)
		assert.Equal(t, 0, enemy.healthPoints)
	})
}
