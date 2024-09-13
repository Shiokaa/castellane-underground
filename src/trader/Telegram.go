package trader

import (
	"fmt"
	"projet-red/inventory"
	"projet-red/object"
)

func Telegram() {
	matraque := object.ObjectStats{"Matraque ", "Arme", 80}
	lacrimogène := object.ObjectStats{"Lacrimogène", "Arme", 15}
	mortier := object.ObjectStats{"Mortier", "Arme", 200}
	taser := object.ObjectStats{"Taser", "Arme", 100}
	ricard := object.ObjectStats{"Ricard", "Soin", 10}
	flash := object.ObjectStats{"Flash", "Soin", 25}
	redbull := object.ObjectStats{"Redbull", "Consumable", 20}

	inv := inventory.Inventory{SacocheCp: []object.ObjectStats{}, Limite: 5}
	achat := 0
	fmt.Println("\nTaper 1 pour acheter une LACRIMOGENE : -5% ‎de point de vie par tour et 1 chance sur 3 de raté son attaque, le tout pendant 3 tours")
	fmt.Println("\nTaper 2 pour acheter un MATRAQUE : Inflige 80 points de dégât")
	fmt.Println("\nTaper 3 pour acheter un MORTIER : Inflige 200 points de dégât")
	fmt.Println("\nTaper 4 pour acheter un TASER : Inflige 100 points de dégât et immobilise l'adversaire pendant 1 tour et utilisable tout les 4 tours ")
	fmt.Println("\nTaper 5 pour acheter un RICARD : Permet de récupérer 10 points de vie")
	fmt.Println("\nTaper 6 pour acheter un FLASH : Permet de récupérer 25 points de vie")
	fmt.Println("\nTaper 7 pour acheter une REDBULL : Permet d'augmenter les dégâts de 20% ‎pendant 3 attaques")
	fmt.Println("\nTaper 8 pour acheter un CASQUE ARAI : Permet d'augmenter son shield de 50% ‎ ")
	fmt.Println("\nTaper 9 pour acheter un ENSEMBLE NIKE TECH : Permet d'augment son shield de 20% ‎ ")
	fmt.Println("\nTaper 10 pour acheter une SACOCHE LV : Rajoute 5 emplacement de plus à l'inventaire ")
	fmt.Println("\nTaper 11 pour sortir du telegram")
	fmt.Scan(&achat)
	for achat >= 0 && achat <= 9 && achat == 10 && achat == 11 {
		fmt.Println("Entrez un numéro valide\n")
		fmt.Scan(&achat)
	}
	if len(inv.SacocheCp) < inv.Limite {
		switch achat {
		case 1:
			inv.SacocheCp = append(inv.SacocheCp, lacrimogène)
			fmt.Println("Une lacrimogène a été ajoutée à l'inventaire")
		case 2:
			inv.SacocheCp = append(inv.SacocheCp, matraque)
			fmt.Print("La matraque a été ajoutée à l'inventaire")
		case 3:
			inv.SacocheCp = append(inv.SacocheCp, mortier)
			fmt.Println("Un mortier a été ajoutée à l'inventaire")
		case 4:
			inv.SacocheCp = append(inv.SacocheCp, taser)
			fmt.Print("Un taser a été ajoutée à l'inventaire")
		case 5:
			inv.SacocheCp = append(inv.SacocheCp, ricard)
			fmt.Println("Un ricard a été ajoutée à l'inventaire")
		case 6:
			inv.SacocheCp = append(inv.SacocheCp, flash)
			fmt.Print("Un flash a été ajoutée à l'inventaire")
		case 7:
			inv.SacocheCp = append(inv.SacocheCp, redbull)
			fmt.Print("Une redbull a été ajoutée à l'inventaire")
		case 8:
			inv.SacocheCp = append(inv.SacocheCp)

			fmt.Println("Vous avez acheté la sacoche lv vous gagnez 10 emplacement")
		default:
		}
	}
}
