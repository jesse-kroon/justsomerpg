package main

type Weapon interface {
	Item
	Damage() int
}

type Sword struct {
	name        string
	value       int
	damage      int
	description string
}

func (s *Sword) Describe() string {
	return s.description
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

func NewSword(name string, value, damage int) *Sword {
	return &Sword{
		name:   name,
		value:  value,
		damage: damage,
	}
}

type Staff struct {
	name        string
	value       int
	damage      int
	description string
}

func (s *Staff) Describe() string {
	return s.description
}

func (s *Staff) Damage() int {
	return s.damage
}

func (s *Staff) Value() int {
	return s.value
}

func (s *Staff) Name() string {
	return s.name
}

func NewStaff(name string, value, damage int) *Staff {
	return &Staff{
		name:   name,
		value:  value,
		damage: damage,
	}
}
