package main

import "fmt"

// CLASS INTRO SCENARIOS
var WarriorIntroScenario = &Scenario{
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
			Enemies:      []*Enemy{NewEnemy(Human, WithLevel(1))},
			End:          "The man lies on the floor, unconscious. You look around to see if there are more attackers. But no one really moves. You greet the innkeeper as you make your way to the establishment's exit.",
			NextScenario: MainChapterOne,
		},
		"Grab the man by his shoulders and tell him to stop, or else...": {},
		"Ignore what's happening and head for the door":                  {},
	},
}

var MageIntroScenario = &Scenario{}

// MAIN SCENARIOS
var MainChapterOne = &Scenario{Dialogue: func() {
	fmt.Println("CONTINUE FROM HERE")
}}
