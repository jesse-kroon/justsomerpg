package main

func newInventory(options ...func(*Inventory)) *Inventory {
	inventory := &Inventory{}

	for _, o := range options {
		o(inventory)
	}

	return inventory
}

func withStartingMoney(amount int) {

}