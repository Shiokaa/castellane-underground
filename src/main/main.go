package main

import (
	"projet-red/game"
	"projet-red/menu"
)

func main() {
	perso := game.ChoixPersonnage()
	inv := game.Firstfight(&perso)
	menu.Menu(&perso, inv)
}
