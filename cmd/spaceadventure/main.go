package main

import (
	"fmt"

	"github.com/ahendricks1/go-spaceadventure-walkthrough/internal/spaceadventure"
)

func main() {
	spaceadventure.PrintWelcome()
	spaceadventure.PrintGreeting(spaceadventure.GetResponseToPrompt("What is your name?"))
	fmt.Println("Let's go on an adventure!")
	spaceadventure.Travel(spaceadventure.GetResponseToPrompt("Shall I randomly choose a planet for you to visit? (Y or N)"))
}
