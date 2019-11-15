package spaceadventure

import "fmt"

func Start(planetarySystem PlanetarySystem) {
	PrintWelcome(planetarySystem)
	PrintGreeting(GetResponseToPrompt("What is your name?"))
	fmt.Println("Let's go on an adventure!")
	Travel(PromptForRandomOrSpecificDestination("Shall I randomly choose a planet for you to visit? (Y or N)"))
}

func PrintWelcome(planetarySystem PlanetarySystem) {
	fmt.Printf("Welcome to the %s!\n", planetarySystem.Name)
	fmt.Printf("There are %d planets to explore.\n", planetarySystem.NumberOfPlanets())
}

func PrintGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

/*
func travel() {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = getResponseToPrompt("Shall I randomly choose a planet for you to visit? (Y or N)")
		if choice == "Y" {
			travelToRandomPlanet()
		} else if choice == "N" {
			planetName := getResponseToPrompt("Name the planet you would like to visit.")
			travelToPlanet(planetName)
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
}
*/

func Travel(randomDestination bool) {
	if (randomDestination) {
		TravelToRandomPlanet()
	} else {
		TravelToPlanet(GetResponseToPrompt("Please name the planet you would like to visit."))
	}
}

func PromptForRandomOrSpecificDestination(prompt string) bool {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = GetResponseToPrompt(prompt)
		if choice == "Y" {
			return true
		} else if choice == "N" {
			return false
		} else {
			fmt.Println("I'm sorry I didn't get that.")
		}
	}
	return false
}

func GetResponseToPrompt(prompt string) string {
	var response string
	fmt.Println(prompt)
	fmt.Scan(&response)
	return response
}

func TravelToPlanet(planetName string) {
	fmt.Printf("Traveling to %s...\n", planetName)
	fmt.Println("Arrived at Neptune. A very cold planet, furthest from the sun.")
}

func TravelToRandomPlanet() {
	fmt.Println("Traveling to Jupiter...")
	fmt.Println("Arrived at Jupiter. The large red spot appears ominous.")
}
