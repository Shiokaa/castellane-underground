package main

import (
	"projet-red/fight"
	"projet-red/game"
	"projet-red/menu"
	"projet-red/story"
)

func main() {
	story.HistoireDebut()
	perso := game.ChoixPersonnage()
	story.Afterchoixperso()
	inv := fight.Firstfight(&perso)
	story.Afterguetteur()
	menu.Menu(&perso, inv)
}
