package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Scenario struct {
	Options      map[int]string
	FollowUp     map[string]*Scenario
	End          string
	NextScenario *Scenario
}

func (s *Scenario) DisplayOptions() {
	for _, option := range s.Options {
		fmt.Println(option)
	}
}

func (s *Scenario) PlayScenario() {
	if len(s.Options) == 0 {
		fmt.Println(s.End)
		return
	}

	playerChoice := getPlayerChoice(s.Options)
	if len(s.FollowUp) != 0 {
		s.FollowUp[playerChoice].DisplayOptions()
		s.FollowUp[playerChoice].PlayScenario()
	}
	fmt.Println(s.End)
	s.NextScenario.PlayScenario()
}

var player *Player

func main() {
	ClearScreen()
	fmt.Println("Welcome to JustSomeRPG. This game is under development.")

	CharacterCreation()
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
	s := &Scenario{
		Options: map[int]string{
			1: "Try to calm the man",
			2: "Grab the man by his shoulders and tell him to stop, or else...",
			3: "Ignore what's happening and head for the door",
		},
		FollowUp: map[string]*Scenario{
			"Try to calm the man": {End: "The man looks at you and ROARSS OMG WHAT?!?!?!"},
			"Grab the man by his shoulders and tell him to stop, or else...": {},
			"Ignore what's happening and head for the door":                  {},
		},
		End:          "Whatever..... TESTINGGGG",
		NextScenario: &Scenario{},
	}
	fmt.Println("You wake up in the inn of Edgewood, a small village in the Southern parts of Kharea...")
	fmt.Println("A vague memory of last night slips your mind, as you try to remember how you got here. You shrug it off, no sense in putting much effort.")
	fmt.Println("On the table in the small room lies your gear, a sword and a shield. You think about how long it's been since you had to use them... Too long for your liking.")
	fmt.Println("You get up from your bed, quickly wash your face and grab your gear. As you head towards the room's exit, you hear a commotion coming from down the stairs.")
	fmt.Println("As you descend the steps, you notice a big, plump man yelling at the innkeeper. His face red, although you can't quite tell if it's from anger or having too much to drink.")
	fmt.Println()

	s.PlayScenario()
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
