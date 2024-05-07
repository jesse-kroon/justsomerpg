package main

import (
	"testing"

	"github.com/jesse-kroon/somerpg/item"
	"github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
	t.Run("should correctly assign health and manapoints based on player class", func(t *testing.T) {
		playerWarrior := newPlayer("", Warrior)
		playerMage := newPlayer("", Mage)

		assert.Equal(t, 0, playerWarrior.manaPoints)
		assert.Equal(t, 15, playerMage.manaPoints)

		assert.Equal(t, 12, playerMage.healthPoints)
		assert.Equal(t, 15, playerWarrior.healthPoints)
	})
}

func TestInventory(t *testing.T) {
	t.Run("characters should start with an empty inventory", func(t *testing.T) {
		playerWarrior := newPlayer("", Warrior)
		playerMage := newPlayer("", Mage)

		assert.Equal(t, playerWarrior.inventory, &Inventory{})
		assert.Equal(t, playerMage.inventory, &Inventory{})
	})

	t.Run("should be able to add an item to inventory", func(t *testing.T) {
		someItem := item.NewItem(1, "stick", "just an ordinary stick")
		player := newPlayer("", Warrior)

		player.addItem(someItem)

		assert.Equal(t, len(player.inventory.items), 1)
		assert.Equal(t, "stick", player.inventory.items[0].Name())
	})
}
