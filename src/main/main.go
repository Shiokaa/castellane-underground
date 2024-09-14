package main

import (
	"fmt"
	"projet-red/game"
	"projet-red/menu"
)

func main() {
	perso := game.ChoixPersonnage()
	inv := game.Firstfight(&perso)
	fmt.Print(perso)
	menu.Menu(&perso, inv)
}
