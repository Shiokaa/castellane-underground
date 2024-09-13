package main

import (
	"projet-red/game"
	"projet-red/menu"
)

func main() {
	perso := game.ChoixPersonnage()
	menu.Menu(perso)
}
