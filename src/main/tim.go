package main

<<<<<<< HEAD
import (
	"projet-red/fight"
	"projet-red/game"
	"projet-red/menu"
	"projet-red/story"
)

func tim() {
	story.HistoireDebut()
	perso := game.ChoixPersonnage()
	story.Afterchoixperso()
	inv := fight.Firstfight(&perso)
	story.Afterguetteur()
	menu.Menu(&perso, inv)
	fight.SecondFight(&perso, inv)
	story.AfterVendeur()
	fight.ThirdFight(&perso, inv)
	story.AfterGoFast()
=======
func tim() {

>>>>>>> 36afc2904dc0cb977f35b9ae041844e9127c7f5d
}
