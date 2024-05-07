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

type Enemy struct {
	name      string
	enemyType EnemyType
	level     int
	weapon    Weapon
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
