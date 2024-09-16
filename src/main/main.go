package main

import (
	"projet-red/game"
	"projet-red/menu"
	"projet-red/story"
)

func main() {

	game.HistoireDebut()
	perso := game.ChoixPersonnage()
	game.Firstfight(perso)

}
