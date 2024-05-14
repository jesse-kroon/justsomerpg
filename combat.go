package main

import (
	"fmt"
	"time"
)

func Combat(p *Player, e *Enemy) {
	time.Sleep(time.Second * 3)

	playerInitiative, EnemyInitiative := rollInitiative()
	playerTurn := playerInitiative > EnemyInitiative

	ClearScreen()
	fmt.Printf("You are now entering combat with %s\n", e.name)
	time.Sleep(time.Second * 1)
	// Enter combat loop
	// TODO: think of a way to add a basic hit/miss mechanic
	for p.healthPoints != 0 && e.healthPoints != 0 {
		if playerTurn {
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("You deal %d damage to the enemy %s\n", p.weapon.Damage(), e.name)
			e.healthPoints -= p.weapon.Damage()
			if e.healthPoints <= 0 {
				fmt.Printf("You have defeated the %s and gained %d experience points!\n", e.name, e.experience)
				p.experiencePoints += e.experience
			}
			playerTurn = false
		} else {
			time.Sleep(time.Millisecond * 500)
			fmt.Printf("The enemy %s deals %d damage to you\n", e.name, e.weapon.Damage())
			p.healthPoints -= e.weapon.Damage()
			if p.healthPoints <= 0 {
				fmt.Printf("You were killed by the %s\n", e.name)
			}
			playerTurn = true
		}
	}
}

// TODO: Randomize this, for now player always starts
func rollInitiative() (int, int) {
	return 2, 1
}
