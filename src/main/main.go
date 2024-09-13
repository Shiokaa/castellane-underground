package main

import (
	"projet-red/game"
)

func main() {
game.HistoireDebut()
perso := game.ChoixPersonnage()
game.Firstfight(perso)
}
