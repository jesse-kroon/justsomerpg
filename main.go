package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var playing bool = true
var player *Player
var playerClassChoice = map[int]Class{
	1: Warrior,
	2: Mage,
}

func main() {
	ClearScreen()
	fmt.Println("Welcome to JustSomeRPG. This game is under development.")

	CharacterCreation()
}

func CharacterCreation() {
	fmt.Println("You are about to create a new character. Please enter a name ...")
	playerName := getPlayerInput()

	ClearScreen()
	fmt.Println("Are you ready to choose a class? For now you can choose between a fearsome Warrior or a spellslinging Mage!")
	playerClass := getPlayerChoice(playerClassChoice)
	player = NewPlayer(playerName, playerClass)

	ClearScreen()
	fmt.Printf("You are %s, level 1 %s. Your journey starts here...\n", player.name, player.class)
	fmt.Println()
	DisplayClassIntro(player.class)
}

func DisplayClassIntro(class Class) {
	switch class {
	case Warrior:
		WarriorIntro()
	case Mage:
		MageIntro()
	}
}

func WarriorIntro() {
	options := map[int]string{
		1: "Try to calm the man",
		2: "Grab the man by his shoulders and tell him to stop, or else...",
		3: "Ignore what's happening and head for the door",
	}
	fmt.Println("You wake up in the inn of Edgewood, a small village in the Southern parts of Kharea...")
	fmt.Println("A vague memory of last night slips your mind, as you try to remember how you got here. You shrug it off, no sense in putting much effort.")
	fmt.Println("On the table in the small room lies your gear, a sword and a shield. You think about how long it's been since you had to use them... Too long for your liking.")
	fmt.Println("You get up from your bed, quickly wash your face and grab your gear. As you head towards the room's exit, you hear a commotion coming from down the stairs.")
	fmt.Println("As you descend the steps, you notice a big, plump man yelling at the innkeeper. His face red, although you can't quite tell if it's from anger or having too much to drink.")
	fmt.Println()
	playerChoice := getPlayerChoice(options)
	fmt.Println(playerChoice)
}

func MageIntro() {

}

func getPlayerInput() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

func getPlayerChoice[T any](options map[int]T) T {
	keys := make([]int, 0, len(options))
	for k := range options {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d.\t%v\n", k, options[k])
	}

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	inputT, _ := strconv.Atoi(input.Text())
	return options[inputT]
}
