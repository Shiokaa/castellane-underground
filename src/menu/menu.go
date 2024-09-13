package menu

import (
	"fmt"
	"projet-red/character"
	"projet-red/inventory"
	"projet-red/object"
	"projet-red/trader"
	"time"
)

func Menu(perso character.Personnage) {
	matraque := object.ObjectStats{"Matraque ", "Arme", 80}
	lacrimogène := object.ObjectStats{"Lacrimogène", "Arme", 15}
	mortier := object.ObjectStats{"Mortier", "Arme", 200}
	taser := object.ObjectStats{"Taser", "Arme", 100}
	ricard := object.ObjectStats{"Ricard", "Soin", 10}
	flash := object.ObjectStats{"Flash", "Soin", 25}
	redbull := object.ObjectStats{"Redbull", "Consumable", 20}

	choix := 0
	inv := inventory.Inventory{SacocheCp: []object.ObjectStats{matraque, lacrimogène, mortier, taser, ricard, flash, redbull}, Limite: 10}

	fmt.Println("\nTapez 1 pour aller au combat		Tapez 2 pour accéder à l'inventaire")
	fmt.Println("Tapez 3 pour aller au O'Tacos		Tapez 4 pour accéder au telegram")
	fmt.Scan(&choix)
	switch choix {
	case 1:
		break
	case 2:
		fmt.Println("Voici l'inventaire :\n")
		for _, v := range inv.SacocheCp {
			fmt.Println(v, "\n")
		}
		Menu(character.Personnage{})
	case 3:
		if perso.Hp < perso.Hpmax {
			fmt.Println("Vous êtes à O'Tacos, vous payez 5€ et régénérer vôtre vie")
			perso.Hp = perso.Hpmax
			time.Sleep(4 * time.Second)
			fmt.Println("Vous avez régénérer tout vos point de vie")
		} else {
			fmt.Println("Vos points de vie sont déjà au maximum")
			Menu(character.Personnage{})
		}
	case 4:
		trader.Telegram()
	}
}
