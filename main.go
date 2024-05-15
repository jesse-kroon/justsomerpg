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

func (s *Scenario) DisplayOptions() {
	for _, option := range s.Options {
		fmt.Println(option)
	}
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

	playerChoice := getPlayerChoice(s.Options)
	if len(s.FollowUp) != 0 {
		s.FollowUp[playerChoice].DisplayOptions()
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
		WarriorIntro()
	case Mage:
		MageIntro()
	}
}

func WarriorIntro() {
	s := &Scenario{
		Dialogue: func() {
			fmt.Println("You wake up in the inn of Edgewood, a small village in the Southern parts of Kharea...")
			fmt.Println("A vague memory of last night slips your mind, as you try to remember how you got here. You shrug it off, no sense in putting much effort.")
			fmt.Println("On the table in the small room lies your gear, a sword and a shield. You think about how long it's been since you had to use them... Too long for your liking.")
			fmt.Println("You get up from your bed, quickly wash your face and grab your gear. As you head towards the room's exit, you hear a commotion coming from down the stairs.")
			fmt.Println("As you descend the steps, you notice a big, plump man yelling at the innkeeper. His face red, although you can't quite tell if it's from anger or having too much to drink.")
			fmt.Println()
		},
		Options: map[int]string{
			1: "Try to calm the man",
			2: "Grab the man by his shoulders and tell him to stop, or else...",
			3: "Ignore what's happening and head for the door",
		},
		FollowUp: map[string]*Scenario{
			"Try to calm the man": {
				Dialogue: func() {
					fmt.Println("The agressive man looks at you and spits you in the face. Flabbergasted, you retreat and hold your fists up.")
				},
				Enemies: []*Enemy{NewEnemy(Human, WithLevel(1))},
				End:     "The man lies on the floor, unconscious. You look around...",
				NextScenario: &Scenario{Dialogue: func() {
					fmt.Println("I haven't actually thought about this...")
				}},
			},
			"Grab the man by his shoulders and tell him to stop, or else...": {},
			"Ignore what's happening and head for the door":                  {},
		},
	}

	s.PlayScenario()
}

func MageIntro() {

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
