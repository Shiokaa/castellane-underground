package menu

import (
	"fmt"
	"projet-red/character"
	"projet-red/game"
	"projet-red/inventory"
	"projet-red/object"
	"time"
)

func Menu(perso *character.Personnage, inv inventory.Inventory) {
	var choix int
	for {
		fmt.Println("\n---------- MENU PRINCIPAL ----------")
		fmt.Println("1 - Aller au combat")
		fmt.Println("2 - Accéder à l'inventaire")
		fmt.Println("3 - Aller au O'Tacos (régénération de vie)")
		fmt.Println("4 - Accéder au Telegram (marché)")
		fmt.Println("5 - Aller zoner dans le quartier")
		fmt.Println("0 - Quitter")
		fmt.Println(" - Vos points de vie ")
		game.DisplayHealth(perso.NameUser, perso.Hp, perso.Hpmax)
		fmt.Println("------------------------------------")
		fmt.Printf("Votre choix: ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			fmt.Println("Vous vous préparez à entrer dans un combat...")
			return
		case 2:
			afficherInventaire(&inv)
		case 3:
			oTacos(perso, inv)
		case 4:
			Telegram(perso, inv)
		case 5:
			fmt.Println("Vous zonez dans le quartier !")
			time.Sleep(15 * time.Second)
			perso.Gold += 25
		case 0:
			fmt.Println("À bientôt!")
			time.Sleep(4 * time.Second)
			return
		default:
			fmt.Println("Choix invalide, essayez à nouveau.")
		}
	}
}

func afficherInventaire(inv *inventory.Inventory) {
	fmt.Println("\n--------- INVENTAIRE ---------")
	if len(inv.SacocheCp) > 0 {
		for i, obj := range inv.SacocheCp {
			fmt.Printf("%d. %s (%s)\n", i+1, obj.Name, obj.Type)
		}
	} else {
		fmt.Println("Votre inventaire est vide.")
	}
	fmt.Println("------------------------------")
	time.Sleep(2 * time.Second)
	fmt.Printf("\nTapez 1 pour supprimer un objet sinon tapez 0. \n\n")
	var choix int
	fmt.Scan(&choix)
	switch choix {
	case 1:
		var Nomobjet string
		fmt.Println("Tapez le nom de l'objet à supprimer")
		fmt.Scan(&Nomobjet)
		for index := range inv.SacocheCp {
			fmt.Println(index, inv.SacocheCp[index])
			if inv.SacocheCp[index].Name == Nomobjet {
				if index == len(inv.SacocheCp)-1 {
					inv.SacocheCp = inv.SacocheCp[:index]
				} else {
					fmt.Println(inv.SacocheCp[:index], inv.SacocheCp[index+1:])
					inv.SacocheCp = append(inv.SacocheCp[:index], inv.SacocheCp[index+1:]...)
					fmt.Println(inv.SacocheCp)
					fmt.Printf("Objet '%s' supprimé avec succès.\n", Nomobjet)
				}
				break
			}
		}
	case 0:
		return
	}
}

func oTacos(perso *character.Personnage, inv inventory.Inventory) {
	clearScreen()
	if perso.Hp < perso.Hpmax {
		if perso.Gold >= 5 {
			perso.Gold -= 5
			perso.Hp = perso.Hpmax
			fmt.Println("Vous êtes allé au O'Tacos, vous avez payé 5€ et vous avez régénéré toute votre vie.")
		} else {
			fmt.Println("Vous n'avez pas assez d'argent pour manger au O'Tacos.")
		}
	} else {
		fmt.Println("Votre vie est déjà au maximum.")
	}
	time.Sleep(3 * time.Second)
	Menu(perso, inv)
}

func Telegram(perso *character.Personnage, inv inventory.Inventory) {
	clearScreen()
	var achat int
	afficherMarché()

	fmt.Printf("\nVous avez %d€.\n", perso.Gold)
	fmt.Println("Entrez le numéro de l'article que vous souhaitez acheter, ou 0 pour quitter:")
	fmt.Scan(&achat)

	switch achat {
	case 1:
		achatObjet(perso, inv, object.ObjectStats{Name: "Lacrimogène", Type: "Arme", Damage: 15}, 30)
	case 2:
		achatObjet(perso, inv, object.ObjectStats{Name: "Matraque", Type: "Arme", Damage: 80}, 100)
	case 3:
		achatObjet(perso, inv, object.ObjectStats{Name: "Mortier", Type: "Arme", Damage: 200}, 150)
	case 4:
		achatObjet(perso, inv, object.ObjectStats{Name: "Taser", Type: "Arme", Damage: 100}, 250)
	case 5:
		achatObjet(perso, inv, object.ObjectStats{Name: "Ricard", Type: "Soin", Damage: 10}, 5)
	case 6:
		achatObjet(perso, inv, object.ObjectStats{Name: "Flash", Type: "Soin", Damage: 25}, 20)
	case 7:
		achatObjet(perso, inv, object.ObjectStats{Name: "Redbull", Type: "Consumable", Damage: 20}, 10)
	case 8:
		if achatStatUpgrade(perso, "Ensemble Nike Tech", 200) {
			perso.Hpmax += 20
			perso.Hp += 20
		}
	case 9:
		if achatStatUpgrade(perso, "Casque Arai", 500) {
			perso.Hpmax += 50
			perso.Hp += 50
		}
	case 10:
		if achatStatUpgrade(perso, "Sacoche LV", 300) {
			inv.Limite += 5
		}
	case 0:
		fmt.Println("Retour au menu principal...")
		time.Sleep(2 * time.Second)
		Menu(perso, inv)
	default:
		fmt.Println("Option invalide.")
		Telegram(perso, inv)
	}
}

func afficherMarché() {
	fmt.Println("\n--------- TELEGRAM MARCHÉ ---------")
	fmt.Println("1 - LACRIMOGENE à 30€ : -5% de vie par tour, 1 chance sur 3 de rater l'attaque (3 tours)")
	fmt.Println("2 - MATRAQUE à 100€ : Inflige 80 dégâts")
	fmt.Println("3 - MORTIER à 150€ : Inflige 200 dégâts")
	fmt.Println("4 - TASER à 250€ : Inflige 100 dégâts et immobilise 1 tour")
	fmt.Println("5 - RICARD à 5€ : Récupère 10 PV")
	fmt.Println("6 - FLASH à 20€ : Récupère 25 PV")
	fmt.Println("7 - REDBULL à 10€ : Augmente dégâts de 20% (3 attaques)")
	fmt.Println("8 - ENSEMBLE NIKE TECH à 200€ : Augmente PV max de 20")
	fmt.Println("9 - CASQUE ARAI à 500€ : Augmente PV max de 50")
	fmt.Println("10 - SACOCHE LV à 300€ : Augmente la capacité d'inventaire de 5 emplacements")
	fmt.Println("0 - Quitter le Telegram")
	fmt.Println("----------------------------------")
}

func achatObjet(perso *character.Personnage, inv inventory.Inventory, objet object.ObjectStats, prix int) {
	if perso.Gold >= prix && len(inv.SacocheCp) < inv.Limite {
		perso.Gold -= prix
		inv.AddObject(objet)
		fmt.Printf("Vous avez acheté %s pour %d€ et l'avez ajouté à l'inventaire.\n", objet.Name, prix)
	} else if perso.Gold < prix {
		fmt.Println("Vous n'avez pas assez d'argent pour cet achat.")
	} else {
		fmt.Println("Votre inventaire est plein.")
	}
	time.Sleep(2 * time.Second)
	Telegram(perso, inv)
}

func achatStatUpgrade(perso *character.Personnage, itemName string, prix int) bool {
	if perso.Gold >= prix {
		perso.Gold -= prix
		fmt.Printf("Vous avez acheté %s et vos statistiques ont été améliorées.\n", itemName)
		time.Sleep(2 * time.Second)
		return true
	} else {
		fmt.Println("Vous n'avez pas assez d'argent.")
		time.Sleep(2 * time.Second)
		return false
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func UseObject(obj object.ObjectStats) int {
	return 0
}
