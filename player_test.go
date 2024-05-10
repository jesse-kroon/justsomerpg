package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
	t.Run("should correctly assign starting health and manapoints based on player class", func(t *testing.T) {
		playerWarrior := NewPlayer("", Warrior)
		playerMage := NewPlayer("", Mage)

		assert.Equal(t, 12, playerMage.healthPoints)
		assert.Equal(t, 15, playerMage.manaPoints)

		assert.Equal(t, 15, playerWarrior.healthPoints)
		assert.Equal(t, 0, playerWarrior.manaPoints)
	})

	t.Run("characters should start with a weapon that is based on their class", func(t *testing.T) {
		playerWarrior := NewPlayer("", Warrior)
		playerMage := NewPlayer("", Mage)

		assert.Equal(t, "Wooden Sword", playerWarrior.weapon.Name())
		assert.Equal(t, "Wooden Staff", playerMage.weapon.Name())
	})

	t.Run("should correctly assign stats to player based on class", func(t *testing.T) {
		playerWarrior := NewPlayer("", Warrior)
		playerMage := NewPlayer("", Mage)

		expectedStatsWarrior := &Attributes{Strength: 9, Stamina: 5, Agility: 7, Toughness: 10, Intellect: 2}
		expectedStatsMage := &Attributes{Strength: 2, Stamina: 5, Agility: 3, Toughness: 6, Intellect: 9}

		assert.Equal(t, expectedStatsMage, playerMage.attributes)
		assert.Equal(t, expectedStatsWarrior, playerWarrior.attributes)
	})
}

func TestPlayer(t *testing.T) {
	t.Run("should be able to gain experience", func(t *testing.T) {
		player := NewPlayer("", Warrior)

		assert.Equal(t, 0, player.experiencePoints)

		player.addExperiencePoints(10)

		assert.Equal(t, 10, player.experiencePoints)
	})

	t.Run("should be able to level up", func(t *testing.T) {
		player := NewPlayer("", Warrior)

		player.addExperiencePoints(player.XPToNextLevel())
		assert.Equal(t, 2, player.level)

		player.addExperiencePoints(player.XPToNextLevel())
		assert.Equal(t, 3, player.level)

		player.addExperiencePoints(player.XPToNextLevel())
		assert.Equal(t, 4, player.level)
	})

	t.Run("should increase MP/HP on level up", func(t *testing.T) {
		playerMage := NewPlayer("", Mage)
		playerWarrior := NewPlayer("", Warrior)

		playerMage.addExperiencePoints(playerMage.XPToNextLevel())
		playerWarrior.addExperiencePoints(playerWarrior.XPToNextLevel())

		assert.Equal(t, 14, playerMage.healthPoints)
		assert.Equal(t, 20, playerMage.manaPoints)

		assert.Equal(t, 20, playerWarrior.healthPoints)
		assert.Equal(t, 0, playerWarrior.manaPoints)
	})
}
