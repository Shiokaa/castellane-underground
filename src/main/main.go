package main

import (
	"projet-red/fight"
	"projet-red/game"
	"projet-red/menu"
	"projet-red/music"
)

func main() {
	go music.Music()
	perso := game.ChoixPersonnage()
	inv := fight.Firstfight(&perso)
	menu.Menu(&perso, inv)
}
