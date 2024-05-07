package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
	t.Run("should correctly assign health and manapoints based on player class", func(t *testing.T) {
		playerWarrior := NewPlayer("", Warrior)
		playerMage := NewPlayer("", Mage)

		assert.Equal(t, 0, playerWarrior.manaPoints)
		assert.Equal(t, 15, playerMage.manaPoints)

		assert.Equal(t, 12, playerMage.healthPoints)
		assert.Equal(t, 15, playerWarrior.healthPoints)
	})

	t.Run("characters should start with a weapon that is based on their class", func(t *testing.T) {
		playerWarrior := NewPlayer("", Warrior)
		playerMage := NewPlayer("", Mage)

		assert.Equal(t, "Wooden Sword", playerWarrior.weapon.Name())
		assert.Equal(t, "Wooden Staff", playerMage.weapon.Name())
	})
}

func TestInventory(t *testing.T) {
	t.Run("characters should start with an empty inventory", func(t *testing.T) {
		playerWarrior := NewPlayer("", Warrior)
		playerMage := NewPlayer("", Mage)

		assert.Equal(t, playerWarrior.inventory, &Inventory{items: []Item{}, value: 0})
		assert.Equal(t, playerMage.inventory, &Inventory{items: []Item{}, value: 0})
	})

	t.Run("should be able to add an item to inventory", func(t *testing.T) {
		someItem := NewItem(1, "stick", "just an ordinary stick")
		player := NewPlayer("", Warrior)

		player.addItem(someItem)

		assert.Equal(t, len(player.inventory.items), 1)
		assert.Equal(t, "stick", player.inventory.items[0].Name())
		assert.Equal(t, 1, player.inventory.value)
	})

	t.Run("should be able to remove an item from player's inventory", func(t *testing.T) {
		someItem := NewItem(1, "stick", "just an ordinary stick")
		player := NewPlayer("", Warrior)

		player.addItem(someItem)

		assert.Equal(t, len(player.inventory.items), 1)
		assert.Equal(t, "stick", player.inventory.items[0].Name())
		assert.Equal(t, 1, player.inventory.value)

		player.removeItem(someItem)

		assert.Equal(t, 0, len(player.inventory.items))
		assert.Equal(t, 0, player.inventory.value)
	})
}

func TestItem(t *testing.T) {
	t.Run("should return correct values for an item", func(t *testing.T) {
		itemName := "Silver Locket"
		itemValue := 10
		itemDescription := "This could have been a family heirloom"
		someItem := NewItem(itemValue, itemName, itemDescription)

		assert.Equal(t, itemValue, someItem.Value())
		assert.Equal(t, itemName, someItem.Name())
		assert.Equal(t, itemDescription, someItem.Describe())
	})

}
