package main

import (
	"projet-red/fight"
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
	inv := fight.Firstfight(&perso)
	story.Afterguetteur()
	menu.Menu(&perso, inv)
>>>>>>> e8deda006bb606e3f6d5bda06c0602fa9a119b91
}
