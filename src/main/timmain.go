package main

import (
	"projet-red/fight"
	"projet-red/game"
	"projet-red/menu"
)

func timmain() {
	//story.HistoireDebut()
	perso := game.ChoixPersonnage()
	//story.Afterchoixperso()
	inv := fight.Firstfight(&perso)
	///story.Afterguetteur()
	menu.Menu(&perso, inv)
	fight.Secondfight(perso)
	//story.AfterVendeur()
	fight.ThirdFight(&perso, inv)
}
