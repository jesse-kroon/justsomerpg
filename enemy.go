package main

import (
	"math/rand"
)

type EnemyType string

const (
	Goblin EnemyType = "Goblin"
	Orc    EnemyType = "Orc"
	Human  EnemyType = "Human"
)

var enemyWeaponPool map[EnemyType]map[int]Weapon = map[EnemyType]map[int]Weapon{
	Goblin: {
		0: &Sword{"Short sword", 1, 2, "A nasty-looking crooked blade, covered in filth"},
		1: &Staff{"Goblin Staff", 1, 3, "A crude staff cut out of dark wood"},
	},
	Orc: {
		0: &Sword{"Jagged Blade", 1, 2, "A rough-looking short blade"},
	},
	Human: {
		0: &Fists{"Bare Knuckles", 0, 1, "Just them good ol' fists"},
	},
}

var enemyBaseExperiencePool map[EnemyType]int = map[EnemyType]int{
	Goblin: 5,
	Orc:    6,
	Human:  3,
}

type Enemy struct {
	name         string
	enemyType    EnemyType
	level        int
	weapon       Weapon
	experience   int
	healthPoints int
}

// This should be used in most cases to create a new enemy
func DefaultNewEnemy(enemyType EnemyType, playerLevel int) *Enemy {
	return NewEnemy(enemyType, WithLevelInPlayerRange(playerLevel))
}

func NewEnemy(enemyType EnemyType, options ...func(*Enemy)) *Enemy {
	enemy := &Enemy{
		name:         string(enemyType),
		enemyType:    enemyType,
		weapon:       randomWeaponForEnemyType(enemyType),
		healthPoints: determineStartingHealthPointsForEnemy(enemyType),
	}

	for _, o := range options {
		o(enemy)
	}

	// Set experience and HP after level is known. Enemy's level could in theory be 0 but this should never occur.
	enemy.experience = enemy.level * enemyBaseExperiencePool[enemyType]
	enemy.healthPoints *= enemy.level
	return enemy
}

func WithLevel(level int) func(*Enemy) {
	return func(e *Enemy) {
		e.level = level
	}
}

func WithLevelInPlayerRange(playerLevel int) func(*Enemy) {
	return func(e *Enemy) {
		min := playerLevel - 2
		max := playerLevel + 2

		if playerLevel-2 <= 0 {
			min = 1
		}

		e.level = rand.Intn(max-min+1) + min
	}
}

func randomWeaponForEnemyType(enemyType EnemyType) Weapon {
	return enemyWeaponPool[enemyType][len(enemyWeaponPool[enemyType])-1]
}

func determineStartingHealthPointsForEnemy(enemyType EnemyType) (baseHP int) {
	switch enemyType {
	case Goblin:
		baseHP = 11
	case Orc:
		baseHP = 13
	default:
		baseHP = 10
	}

	return
}
