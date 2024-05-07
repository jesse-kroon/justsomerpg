package main

func NewInventory() *Inventory {
	return &Inventory{items: []Item{}, value: 0}
}
