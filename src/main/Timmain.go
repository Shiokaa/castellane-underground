package main

import (
	"projet-red/fight"
	"projet-red/game"
	"projet-red/menu"
	"projet-red/story"
)

func Tomymain() {
	story.HistoireDebut()
	perso := game.ChoixPersonnage()
	story.Afterchoixperso()
	inv := fight.Firstfight(&perso)
	story.Afterguetteur()
	menu.Menu(&perso, inv)
	story.AfterVendeur()
	fight.ThirdFight(&perso, inv)
}
