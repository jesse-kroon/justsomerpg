package main

import "fmt"

func Combat(p *Player, e *Enemy) {
	playerInitiative, EnemyInitiative := rollInitiative()
	playerTurn := playerInitiative > EnemyInitiative

	// Enter combat loop
	// TODO: think of a way to add a basic hit/miss mechanic
	for p.healthPoints != 0 && e.healthPoints != 0 {
		if playerTurn {
			fmt.Printf("You deal %d damage to the enemy %s\n", p.weapon.Damage(), e.name)
			e.healthPoints -= p.weapon.Damage()
			if e.healthPoints <= 0 {
				fmt.Printf("You have defeated the %s and gained %d experience points!\n", e.name, e.experience)
				p.experiencePoints += e.experience
			}
			playerTurn = false
		} else {
			fmt.Printf("The enemy %s deals %d damage to you\n", e.name, e.weapon.Damage())
			p.healthPoints -= e.weapon.Damage()
			if p.healthPoints <= 0 {
				fmt.Printf("You were killed by the %s\n", e.name)
			}
			playerTurn = true
		}
	}
}

// Randomize this, for now player always starts
func rollInitiative() (int, int) {
	return 2, 1
}
