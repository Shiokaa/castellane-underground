package menu

import (
	"fmt"
	"projet-red/character"
	"projet-red/inventory"
	"projet-red/object"
	"time"
)

func Menu(perso *character.Personnage, inv inventory.Inventory) {
	choix := 0
	fmt.Println("\nTapez 1 pour aller au combat		Tapez 2 pour accéder à l'inventaire")
	fmt.Println("Tapez 3 pour aller au O'Tacos		Tapez 4 pour accéder au telegram")
	fmt.Scan(&choix)
	switch choix {
	case 1:
		break
	case 2:
		fmt.Println("Voici l'inventaire :\n")
		fmt.Println(inv.SacocheCp)
		Menu(perso, inv)
	case 3:
		if perso.Hp < perso.Hpmax {
			if perso.Gold >= 5 {
				fmt.Println("Vous êtes à O'Tacos, vous payez 5€ et régénérer vôtre vie")
				perso.Hp = perso.Hpmax
				time.Sleep(4 * time.Second)
				fmt.Println("Vous avez régénérer tout vos point de vie")
				Menu(perso, inv)
			} else {
				fmt.Println("Vous êtes trop pauvre !")
				Telegram(perso, inv)
			}
		} else {
			fmt.Println("Vos points de vie sont déjà au maximum")
			Menu(perso, inv)
		}
	case 4:
		Telegram(perso, inv)
	}
}

func Telegram(perso *character.Personnage, inv inventory.Inventory) {
	matraque := object.ObjectStats{"Matraque ", "Arme", 80}
	lacrimogène := object.ObjectStats{"Lacrimogène", "Arme", 15}
	mortier := object.ObjectStats{"Mortier", "Arme", 200}
	taser := object.ObjectStats{"Taser", "Arme", 100}
	ricard := object.ObjectStats{"Ricard", "Soin", 10}
	flash := object.ObjectStats{"Flash", "Soin", 25}
	redbull := object.ObjectStats{"Redbull", "Consumable", 20}
	achat := 0
	fmt.Println("\nTaper 1 pour acheter une LACRIMOGENE à 30€: -5% ‎de point de vie par tour et 1 chance sur 3 de raté son attaque, le tout pendant 3 tours")
	fmt.Println("\nTaper 2 pour acheter un MATRAQUE à 100€ : Inflige 80 points de dégât")
	fmt.Println("\nTaper 3 pour acheter un MORTIER à 150€ : Inflige 200 points de dégât")
	fmt.Println("\nTaper 4 pour acheter un TASER à 250€ : Inflige 100 points de dégât et immobilise l'adversaire pendant 1 tour et utilisable tout les 4 tours ")
	fmt.Println("\nTaper 5 pour acheter un RICARD 5€ : Permet de récupérer 10 points de vie")
	fmt.Println("\nTaper 6 pour acheter un FLASH 20€ : Permet de récupérer 25 points de vie")
	fmt.Println("\nTaper 7 pour acheter une REDBULL 10€ : Permet d'augmenter les dégâts de 20% ‎pendant 3 attaques")
	fmt.Println("\nTaper 8 pour acheter un ENSEMBLE NIKE TECH 200€ : Permet d'augmenter ses points de vie de 20")
	fmt.Println("\nTaper 9 pour acheter un CASQUE ARAI 500€ : Permet d'augmenter ses points de vie de 50")
	fmt.Println("\nTaper 10 pour acheter une SACOCHE LV 300€ : Rajoute 5 emplacement de plus à l'inventaire")
	fmt.Println("\nTaper 11 pour sortir du telegram")
	fmt.Println("\nVous avez", perso.Gold, "€")
	fmt.Scan(&achat)
	for achat >= 0 && achat <= 9 && achat == 10 && achat == 11 {
		fmt.Println("Entrez un numéro valide\n")
		fmt.Scan(&achat)
	}
	if len(inv.SacocheCp) < inv.Limite {
		switch achat {
		case 1:
			if perso.Gold >= 30 {
				perso.Gold -= 30
				inv.AddObject(lacrimogène)
				fmt.Println("Une lacrimogène a été ajoutée à l'inventaire")
				Menu(perso, inv)
			} else {
				fmt.Println("Vous êtes trop pauvre !")
				Menu(perso, inv)
			}
		case 2:
			if perso.Gold >= 100 {
				perso.Gold -= 100
				inv.AddObject(matraque)
				fmt.Print("La matraque a été ajoutée à l'inventaire")
				Menu(perso, inv)
			} else {
				fmt.Println("Vous êtes trop pauvre !")
				Menu(perso, inv)
			}
		case 3:
			if perso.Gold >= 150 {
				perso.Gold -= 150
				inv.AddObject(mortier)
				fmt.Println("Un mortier a été ajoutée à l'inventaire")
				Menu(perso, inv)
			} else {
				fmt.Println("Vous êtes trop pauvre !")
				Menu(perso, inv)
			}
		case 4:
			if perso.Gold >= 250 {
				perso.Gold -= 250
				inv.AddObject(taser)
				fmt.Print("Un taser a été ajoutée à l'inventaire")
				Menu(perso, inv)
			} else {
				fmt.Println("Vous êtes trop pauvre !")
				Menu(perso, inv)
			}
		case 5:
			if perso.Gold >= 5 {
				perso.Gold -= 5
				inv.AddObject(ricard)
				fmt.Println("Un ricard a été ajoutée à l'inventaire")
				Menu(perso, inv)
			} else {
				fmt.Println("Vous êtes trop pauvre !")
				Menu(perso, inv)
			}
		case 6:
			if perso.Gold >= 25 {
				perso.Gold -= 25
				inv.AddObject(flash)
				fmt.Print("Un flash a été ajoutée à l'inventaire")
				Menu(perso, inv)
			} else {
				fmt.Println("Vous êtes trop pauvre !")
				Menu(perso, inv)
			}
		case 7:
			if perso.Gold >= 10 {
				perso.Gold -= 10
				inv.AddObject(redbull)
				fmt.Print("Une redbull a été ajoutée à l'inventaire")
				Menu(perso, inv)
			} else {
				fmt.Println("Vous êtes trop pauvre !")
				Menu(perso, inv)
			}
		case 8:
			if perso.Gold >= 200 {
				perso.Gold -= 200
				perso.Hpmax += 20
				perso.Hp += 20
				fmt.Print("Vous équipez votre ensemble")
				Menu(perso, inv)
			} else {
				fmt.Println("Vous êtes trop pauvre !")
				Menu(perso, inv)
			}
		case 9:
			if perso.Gold >= 500 {
				perso.Gold -= 500
				perso.Hpmax += 50
				perso.Hp += 50
				fmt.Print("Vous équipez votre casque")
				Menu(perso, inv)
			} else {
				fmt.Println("Vous êtes trop pauvre !")
				Menu(perso, inv)
			}
		case 10:
			if perso.Gold > 300 {
				inv.Limite += 5
				fmt.Println("Vous avez acheté la sacoche lv vous gagnez 10 emplacement")
				Menu(perso, inv)
			} else {
				fmt.Println("Vous êtes trop pauvre !")
				Menu(perso, inv)
			}
		case 11:
			Menu(perso, inv)
		default:
		}
	} else {
		fmt.Println("Inventaire plein ! Impossible d'ajouter plus d'objets.")
	}
}
