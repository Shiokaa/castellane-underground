package main

import (
	"projet-red/fight"
	"projet-red/game"
	"projet-red/menu"
	"projet-red/music"
	"projet-red/story"
)

func tim() {
	go music.Music()
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
}
