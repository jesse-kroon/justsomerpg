package main

import (
	"slices"
)

type Class string

const (
	Warrior Class = "warrior"
	Mage    Class = "mage"

	BASE_HEALTHPOINTS = 10
)

type Inventory struct {
	items    []Item
	value    int
	currency int
}

type Player struct {
	name         string
	healthPoints int
	manaPoints   int
	class        Class
	inventory    *Inventory
	weapon       Weapon
}

func (p *Player) addItem(item Item) {
	p.inventory.items = append(p.inventory.items, item)
	p.inventory.value += item.Value()
}

func (p *Player) removeItem(itemToDelete Item) {
	p.inventory.items = slices.DeleteFunc(p.inventory.items, func(item Item) bool {
		return item.Name() == itemToDelete.Name()
	})
	p.inventory.value -= itemToDelete.Value()
}

func NewPlayer(name string, class Class) *Player {
	return &Player{
		name:         name,
		class:        class,
		manaPoints:   determineManaPoints(class),
		healthPoints: determineHealthPoints(class),
		weapon:       newWeapon(class),
		inventory:    newInventory(),
	}
}

func newWeapon(class Class) Weapon {
	switch class {
	case Warrior:
		return NewSword("Wooden Sword", 2, 2)
	case Mage:
		return NewStaff("Wooden Staff", 2, 3)
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
