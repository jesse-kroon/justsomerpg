package main

import "github.com/jesse-kroon/somerpg/item"

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
}

type Player struct {
	name         string
	healthPoints int
	manaPoints   int
	class        Class
	inventory    *Inventory
}

func newInventory() *Inventory {
	return &Inventory{}
}

func (p *Player) addItem(item item.Item) {
	p.inventory.items = append(p.inventory.items, item)
}

func newPlayer(name string, class Class) *Player {
	return &Player{
		name:         name,
		class:        class,
		manaPoints:   determineManaPoints(class),
		healthPoints: determineHealthPoints(class),
		inventory:    newInventory(),
	}
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
