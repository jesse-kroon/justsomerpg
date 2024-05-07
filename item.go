package main

type Item interface {
	Value() int
	Name() string
	Describe() string
}

type CommonItem struct {
	name        string
	description string
	value       int
}

func (c *CommonItem) Value() int {
	return c.value
}

func (c *CommonItem) Name() string {
	return c.name
}

func (c *CommonItem) Describe() string {
	return c.description
}

func NewItem(value int, name, description string) Item {
	return &CommonItem{
		name,
		description,
		value,
	}
}
