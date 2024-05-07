package item

type Weapon interface {
	Damage() int
}

type Sword struct {
	name        string
	value       int
	damage      int
	description string
}

func (s *Sword) Damage() int {
	return s.damage
}

func (s *Sword) Value() int {
	return s.value
}

func (s *Sword) Name() string {
	return s.name
}

func (s *Sword) Description() string {
	return s.description
}

func NewWeapon(name, category string, value, damage int) Weapon {
	switch category {
	case "sword":
		return &Sword{
			name:   name,
			value:  value,
			damage: damage,
		}
	}

	return &Sword{
		name:   "Broken Sword",
		value:  2,
		damage: 5,
	}
}
