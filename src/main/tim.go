package main

import (
	"projet-red/fight"
	"projet-red/game"
	"projet-red/menu"
	"projet-red/story"
)

func tim() {
	story.HistoireDebut()
	perso := game.ChoixPersonnage()
	inv := fight.Firstfight(&perso)
	menu.Menu(&perso, inv)
}
