package main

import (
	"projet-red/game"
	"projet-red/menu"
	"projet-red/story"
)

func main() {
<<<<<<< HEAD

	game.HistoireDebut()
	perso := game.ChoixPersonnage()
	game.Firstfight(perso)
=======
	story.HistoireDebut()
	perso := game.ChoixPersonnage()
	story.Afterchoixperso()
	game.Firstfight(perso)
	story.Afterguetteur()
	menu.Menu(perso)
>>>>>>> f13e4f7e053c42328883fddd5af8e8285a81a77c

}
