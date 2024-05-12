package main

import (
	"bufio"
	"fmt"
	"os"
)

var playing bool = true
var player *Player
var playerClassChoice = map[string]Class{
	"1": Warrior,
	"2": Mage,
}

func main() {
	ClearScreen()
	fmt.Println("Welcome to JustSomeRPG. This game is under development.")
	fmt.Println("You are about to create a new character. Please enter a name ...")

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	playerName := input.Text()

	fmt.Println("Are you ready to choose a class? For now you can choose between a fearsome Warrior or a spellslinging Mage!")
	fmt.Println("1.\tWarrior")
	fmt.Println("2.\tMage")
	input.Scan()
	playerChoice := input.Text()
	if playerClass, ok := playerClassChoice[playerChoice]; ok {
		player = NewPlayer(playerName, playerClass)
	} else {
		panic("Unsupported class. Shutting down for now....")
	}

	fmt.Printf("You are %s, level 1 %s. Your journey starts here...\n", player.name, player.class)
}
