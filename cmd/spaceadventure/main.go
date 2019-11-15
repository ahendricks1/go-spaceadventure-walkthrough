package main

import "github.com/ahendricks1/go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	var ps = spaceadventure.PlanetarySystem{Name: "Solare System", Planets: []spaceadventure.Planet{

		spaceadventure.Planet{Name: "Tatooine", Description: "Test"},
	} }
	spaceadventure.Start(ps)
}
