package main

func tim( ){
	go.music.Music()
	story.HistoireDebut()
	perso := game.ChoixPersonnage()
	story.Afterchoixperso()
	inv := fight.Firstfight(&perso)
	story.Afterguetteur()
	menu.Menu(&perso, inv)
	fight.SecondFight(perso)
	story.AfterVendeur()
	fight.ThirdFight(&perso, inv)
	story.AfterGoFast()
}