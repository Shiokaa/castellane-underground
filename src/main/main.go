package main

import (
	"projet-red/game"
	"projet-red/menu"
	"projet-red/story"
)

func main() {
	story.HistoireDebut()
	perso := game.ChoixPersonnage()
	story.Afterchoixperso()
	game.Firstfight(perso)
	story.Afterguetteur()
	menu.Menu(perso)

}
