package menu

import (
	"fmt"
	"projet-red/character"
	"projet-red/inventory"
	"projet-red/object"
	"time"
)

// Fonction principale du menu
func Menu(perso *character.Personnage, inv inventory.Inventory) {
	var choix int
	clearScreen()

	for {
		fmt.Println("\n---------- MENU PRINCIPAL ----------")
		fmt.Println("1 - Aller au combat")
		fmt.Println("2 - Accéder à l'inventaire")
		fmt.Println("3 - Aller au O'Tacos (régénération de vie)")
		fmt.Println("4 - Accéder au Telegram (marché)")
		fmt.Println("0 - Quitter")
		fmt.Println("------------------------------------")
		fmt.Printf("Votre choix: ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			fmt.Println("Vous vous préparez à entrer dans un combat...")
			return // Quitte le menu pour démarrer le combat
		case 2:
			afficherInventaire(inv)
		case 3:
			oTacos(perso, inv)
		case 4:
			Telegram(perso, inv)
		case 0:
			fmt.Println("À bientôt!")
			break
		default:
			fmt.Println("Choix invalide, essayez à nouveau.")
		}
	}
}

// Fonction d'affichage de l'inventaire
func afficherInventaire(inv inventory.Inventory) {
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
}

// Fonction pour aller au O'Tacos
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

// Fonction du Telegram (marché)
func Telegram(perso *character.Personnage, inv inventory.Inventory) {
	clearScreen()
	var achat int
	afficherMarché()

	fmt.Printf("\nVous avez %d€.\n", perso.Gold)
	fmt.Println("Entrez le numéro de l'article que vous souhaitez acheter, ou 0 pour quitter:")
	fmt.Scan(&achat)

	// Gestion des achats
	switch achat {
	case 1:
		achatObjet(perso, inv, object.ObjectStats{"Lacrimogène", "Arme", 15}, 30)
	case 2:
		achatObjet(perso, inv, object.ObjectStats{"Matraque", "Arme", 80}, 100)
	case 3:
		achatObjet(perso, inv, object.ObjectStats{"Mortier", "Arme", 200}, 150)
	case 4:
		achatObjet(perso, inv, object.ObjectStats{"Taser", "Arme", 100}, 250)
	case 5:
		achatObjet(perso, inv, object.ObjectStats{"Ricard", "Soin", 10}, 5)
	case 6:
		achatObjet(perso, inv, object.ObjectStats{"Flash", "Soin", 25}, 20)
	case 7:
		achatObjet(perso, inv, object.ObjectStats{"Redbull", "Consumable", 20}, 10)
	case 8:
		if achatStatUpgrade(perso, inv, "Ensemble Nike Tech", 200, 20) {
			perso.Hpmax += 20
			perso.Hp += 20
		}
	case 9:
		if achatStatUpgrade(perso, inv, "Casque Arai", 500, 50) {
			perso.Hpmax += 50
			perso.Hp += 50
		}
	case 10:
		if achatStatUpgrade(perso, inv, "Sacoche LV", 300, 5) {
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

// Fonction pour afficher le marché du Telegram
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

// Fonction générique pour acheter un objet
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

// Fonction générique pour améliorer les stats (ex: HP max, capacité d'inventaire)
func achatStatUpgrade(perso *character.Personnage, inv inventory.Inventory, itemName string, prix int, upgradeValue int) bool {
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

// Fonction pour simuler le nettoyage de l'écran (utile pour lisibilité)
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
