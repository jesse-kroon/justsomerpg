package main

import (
	"slices"

	"github.com/jesse-kroon/somerpg/item"
)

type Class string
type ItemCategory string

const (
	Warrior Class = "warrior"
	Mage    Class = "mage"

	Weapon ItemCategory = "weapon"

	BASE_HEALTHPOINTS = 10
)

type Inventory struct {
	items []item.Item
	value int
}

type Player struct {
	name         string
	healthPoints int
	manaPoints   int
	class        Class
	inventory    *Inventory
	weapon       item.Weapon
}

func newInventory() *Inventory {
	return &Inventory{items: []item.Item{}, value: 0}
}

func (p *Player) addItem(item item.Item) {
	p.inventory.items = append(p.inventory.items, item)
	p.inventory.value += item.Value()
}

func (p *Player) removeItem(itemToDelete item.Item) {
	p.inventory.items = slices.DeleteFunc(p.inventory.items, func(item item.Item) bool {
		return item.Name() == itemToDelete.Name()
	})
	p.inventory.value -= itemToDelete.Value()
}

func newPlayer(name string, class Class) *Player {
	return &Player{
		name:         name,
		class:        class,
		manaPoints:   determineManaPoints(class),
		healthPoints: determineHealthPoints(class),
		weapon:       newWeapon(class),
		inventory:    newInventory(),
	}
}

func newWeapon(class Class) item.Weapon {
	switch class {
	case Warrior:
		return item.NewSword("Wooden Sword", 2, 2)
	case Mage:
		return item.NewStaff("Wooden Staff", 2, 3)
	}

	return nil
}

func determineManaPoints(class Class) int {
	var manaPoints int

	switch class {
	case Warrior:
		manaPoints = 0
	case Mage:
		manaPoints = 15
	}

	return manaPoints
}

func determineHealthPoints(class Class) int {
	var additionalHealthPoints int = 0

	switch class {
	case Warrior:
		additionalHealthPoints = 5
	case Mage:
		additionalHealthPoints = 2
	}

	return BASE_HEALTHPOINTS + additionalHealthPoints
}
