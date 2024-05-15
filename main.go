package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"time"
)

type Scenario struct {
	Dialogue     func()
	Options      map[int]string
	FollowUp     map[string]*Scenario
	End          string
	NextScenario *Scenario
	Enemies      []*Enemy
}

func (s *Scenario) PlayScenario() {
	s.Dialogue()

	// Check if there's anything to do in this scenario
	if len(s.Options) == 0 && len(s.Enemies) == 0 {
		fmt.Println(s.End)
		return
	}

	if len(s.Enemies) != 0 {
		for _, enemy := range s.Enemies {
			Combat(player, enemy)
		}
	}

	if len(s.FollowUp) != 0 {
		playerChoice := getPlayerChoice(s.Options)
		s.FollowUp[playerChoice].PlayScenario()
	}
	fmt.Println(s.End)

	// Remove this at one point
	if s.NextScenario != nil {
		time.Sleep(time.Second * 2)
		ClearScreen()
		s.NextScenario.PlayScenario()
	} else {
		fmt.Println("You have reached the end for now")
	}
}

var player *Player

func main() {
	ClearScreen()
	fmt.Println("Welcome to JustSomeRPG. This game is under development.")

	CharacterCreation()
	StartGame(player.class)
}

func CharacterCreation() {
	playerClassChoice := map[int]Class{
		1: Warrior,
		2: Mage,
	}

	fmt.Println("You are about to create a new character. Please enter a name ...")
	playerName := getPlayerInput()

	ClearScreen()
	fmt.Println("Are you ready to choose a class? For now you can choose between a fearsome Warrior or a spellslinging Mage!")
	playerClass := getPlayerChoice(playerClassChoice)
	player = NewPlayer(playerName, playerClass)

	ClearScreen()
	fmt.Printf("You are %s, level 1 %s. Your journey starts here...\n", player.name, player.class)
	fmt.Println()
}

func StartGame(class Class) {
	switch class {
	case Warrior:
		WarriorIntroScenario.PlayScenario()
	case Mage:
		MageIntroScenario.PlayScenario()
	}
}

func getPlayerInput() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

// TODO: Think of ways to make this function smaller
func getPlayerChoice[T any](options map[int]T) T {
	var err error
	var chosenInput int
	var validInput bool = false

	keys := make([]int, 0, len(options))
	for k := range options {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d.\t%v\n", k, options[k])
	}

	for !validInput {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		chosenInput, err = strconv.Atoi(input.Text())
		if err != nil {
			fmt.Println("Please enter a number corresponding to the option you would like to perform")
			time.Sleep(time.Second * 1)
		} else {
			validInput = true
		}
	}

	validOption := slices.Contains(keys, chosenInput)
	for !validOption {
		fmt.Println("You did not enter a valid option. Please try again...")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		chosenInput, err = strconv.Atoi(input.Text())
		if err != nil {
			fmt.Println("Please enter a number corresponding to the option you would like to perform")
			time.Sleep(time.Millisecond * 500)
		}

		if slices.Contains(keys, chosenInput) {
			validOption = true
			break
		}
	}

	return options[chosenInput]
}
