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

	t.Run("characters should start with a weapon that is based on their class", func(t *testing.T) {
		playerWarrior := newPlayer("", Warrior)
		playerMage := newPlayer("", Mage)

		assert.Equal(t, "Wooden Sword", playerWarrior.weapon.Name())
		assert.Equal(t, "Wooden Staff", playerMage.weapon.Name())
	})
}

func TestInventory(t *testing.T) {
	t.Run("characters should start with an empty inventory", func(t *testing.T) {
		playerWarrior := newPlayer("", Warrior)
		playerMage := newPlayer("", Mage)

		assert.Equal(t, playerWarrior.inventory, &Inventory{items: []item.Item{}, value: 0})
		assert.Equal(t, playerMage.inventory, &Inventory{items: []item.Item{}, value: 0})
	})

	t.Run("should be able to add an item to inventory", func(t *testing.T) {
		someItem := item.NewItem(1, "stick", "just an ordinary stick")
		player := newPlayer("", Warrior)

		player.addItem(someItem)

		assert.Equal(t, len(player.inventory.items), 1)
		assert.Equal(t, "stick", player.inventory.items[0].Name())
		assert.Equal(t, 1, player.inventory.value)
	})

	t.Run("should be able to remove an item from player's inventory", func(t *testing.T) {
		someItem := item.NewItem(1, "stick", "just an ordinary stick")
		player := newPlayer("", Warrior)

		player.addItem(someItem)

		assert.Equal(t, len(player.inventory.items), 1)
		assert.Equal(t, "stick", player.inventory.items[0].Name())
		assert.Equal(t, 1, player.inventory.value)

		player.removeItem(someItem)

		assert.Equal(t, 0, len(player.inventory.items))
		assert.Equal(t, 0, player.inventory.value)
	})
}

func TestWeapon(t *testing.T) {
}

func TestItem(t *testing.T) {
	t.Run("should be able to get the value of an item", func(t *testing.T) {
		someItem := item.NewItem(10, "Silver Locket", "this could have been a family heirloom")

		assert.Equal(t, 10, someItem.Value())
	})
}
