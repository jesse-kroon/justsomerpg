package main

import "math/rand"

type EnemyType string

const (
	Goblin EnemyType = "Goblin"
	Orc    EnemyType = "Orc"
)

var enemyWeaponPool map[EnemyType]map[int]Weapon = map[EnemyType]map[int]Weapon{
	Goblin: {
		0: &Sword{"Short sword", 1, 2, "A rough-looking short blade"},
		1: &Staff{"Goblin Staff", 1, 3, "A crude staff cut out of dark wood"},
	},
	Orc: {
		0: &Sword{"Jagged Blade", 1, 2, "A rough-looking short blade"},
	},
}

var enemyBaseExperiencePool map[EnemyType]float64 = map[EnemyType]float64{
	Goblin: 5,
	Orc:    6,
}

type Enemy struct {
	name       string
	enemyType  EnemyType
	level      int
	weapon     Weapon
	experience float64
}

// This should be used in most cases to create a new enemy
func DefaultNewEnemy(enemyType EnemyType, playerLevel int) *Enemy {
	return NewEnemy(enemyType, WithLevelInPlayerRange(playerLevel))
}

func NewEnemy(enemyType EnemyType, options ...func(*Enemy)) *Enemy {
	enemy := &Enemy{
		name:      string(enemyType),
		enemyType: enemyType,
		weapon:    randomWeaponForEnemyType(enemyType),
	}

	for _, o := range options {
		o(enemy)
	}

	// Set experience after level is known. Enemy's level could in theory be 0 but this should never occur.
	enemy.experience = float64(enemy.level) * enemyBaseExperiencePool[enemyType]

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
