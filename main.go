package main

import "fmt"

func main() {
	player := NewPlayer("Mario", Mage)

	fmt.Println(fmt.Sprintf("%v", player))
}
