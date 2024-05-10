package main

import (
	"math"
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
	name             string
	level            int
	healthPoints     int
	manaPoints       int
	class            Class
	inventory        *Inventory
	weapon           Weapon
	experiencePoints int
}

func NewPlayer(name string, class Class, options ...func(*Player)) *Player {
	player := &Player{
		name:             name,
		class:            class,
		level:            1,
		experiencePoints: 0,
		manaPoints:       determineStartingManaPoints(class),
		healthPoints:     determineStartingHealthPoints(class),
		weapon:           determineStartingWeapon(class),
		inventory:        NewInventory(),
	}

	for _, o := range options {
		o(player)
	}

	return player
}

func WithStartingCurrency(amount int) func(*Player) {
	return func(p *Player) {
		p.inventory.currency = amount
	}
}

func WithClassBasedStartingInventory() func(*Player) {
	return func(p *Player) {
		switch p.class {
		case Warrior:
			p.inventory.items = []Item{
				&CommonItem{name: "Whetstone", value: 1, description: "You can use this item to sharpen your sword"},
				&CommonItem{name: "Torch", value: 1, description: "Use this to light your way in the darkest places of this world..."},
			}
		case Mage:
			p.inventory.items = []Item{
				&CommonItem{name: "Magic Orb", value: 1, description: "A memento from your time at the mages guild."},
				&CommonItem{name: "Torch", value: 1, description: "Use this to light your way in the darkest places of this world..."},
			}
		}
	}
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

func (p *Player) addExperiencePoints(amount int) {
	p.experiencePoints += amount

	// Check if experience points are >= required for level up, if so, level up
	if p.experiencePoints >= ExperienceRequiredForNextLevel(p.level) {
		p.level++
	}
}

func (p *Player) XPToNextLevel() int {
	return ExperienceRequiredForNextLevel(p.level) - p.experiencePoints
}

func ExperienceRequiredForNextLevel(currentLevel int) int {
	baseXP := 500
	return int(math.Floor(float64(baseXP) * math.Pow(float64(currentLevel), 1.5)))
}

func determineStartingWeapon(class Class) Weapon {
	switch class {
	case Warrior:
		return NewSword("Wooden Sword", 2, 2)
	case Mage:
		return NewStaff("Wooden Staff", 2, 3)
	}

	return nil
}

func determineStartingManaPoints(class Class) int {
	var manaPoints int

	switch class {
	case Warrior:
		manaPoints = 0
	case Mage:
		manaPoints = 15
	}

	return manaPoints
}

func determineStartingHealthPoints(class Class) int {
	var additionalHealthPoints int = 0

	switch class {
	case Warrior:
		additionalHealthPoints = 5
	case Mage:
		additionalHealthPoints = 2
	}

	return BASE_HEALTHPOINTS + additionalHealthPoints
}
