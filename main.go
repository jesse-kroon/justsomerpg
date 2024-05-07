package main

import "fmt"

func main() {
	player := newPlayer("Mario", Mage)

	fmt.Println(fmt.Sprintf("%v", player))
}
