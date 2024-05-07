package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEnemy(t *testing.T) {
	t.Run("should be able to create an enemy with specified level", func(t *testing.T) {
		enemy := NewEnemy("Goblin", WithLevel(2))

		assert.Equal(t, 2, enemy.level)
	})

	t.Run("should be able to create an enemy within player range", func(t *testing.T) {
		player := NewPlayer("", Warrior)
		player.level = 10
		enemy := NewEnemy("Goblin", WithLevelInPlayerRange(player.level))

		assert.Positive(t, enemy.level)
		assert.True(t, enemy.level <= player.level+2 && enemy.level >= player.level-2)
	})

	t.Run("should create an enemy with a weapon based on the enemy type", func(t *testing.T) {
		enemy := NewEnemy(Goblin)
		correctWeapon := false

		// Look in the Goblin's weapon pool
		for _, v := range enemyWeaponPool[Goblin] {
			if v.Name() == enemy.weapon.Name() {
				correctWeapon = true
			}
		}
		assert.True(t, correctWeapon)

		// Now look in the Orc's weapon pool
		correctWeapon = false
		for _, v := range enemyWeaponPool[Orc] {
			if v.Name() == enemy.weapon.Name() {
				correctWeapon = true
			}
		}
		assert.False(t, correctWeapon)
	})

}
