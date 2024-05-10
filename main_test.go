package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestInventory(t *testing.T) {
	assertInventoryContains := func(t testing.TB, inventory []Item, want string) {
		if !slices.ContainsFunc(inventory, func(item Item) bool {
			return item.Name() == want
		}) {
			t.Errorf("expected inventory to hold a '%s'", want)
		}
	}

	t.Run("characters should start with an empty inventory when no starting inventory is configured", func(t *testing.T) {
		playerWarrior := NewPlayer("", Warrior)
		playerMage := NewPlayer("", Mage)

		assert.Equal(t, playerWarrior.inventory, &Inventory{items: []Item{}, value: 0})
		assert.Equal(t, playerMage.inventory, &Inventory{items: []Item{}, value: 0})
	})

	t.Run("inventory should hold currency when created with starting currency", func(t *testing.T) {
		player := NewPlayer("", Warrior, WithStartingCurrency(10))

		assert.Equal(t, 10, player.inventory.currency)
	})

	t.Run("should be able to start a character with default inventory based on class", func(t *testing.T) {
		player := NewPlayer("", Warrior, WithClassBasedStartingInventory())
		assertInventoryContains(t, player.inventory.items, "Whetstone")

		player = NewPlayer("", Mage, WithClassBasedStartingInventory())
		assertInventoryContains(t, player.inventory.items, "Magic Orb")
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
