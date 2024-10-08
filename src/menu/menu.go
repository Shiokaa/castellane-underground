package menu

import (
	"fmt"
	"os"
	"projet-red/ascii"
	"projet-red/character"
	"projet-red/fight"
	"projet-red/game"
	"projet-red/inventory"
	"projet-red/jeutel"
	"projet-red/music"
	"projet-red/object"
	"time"
)

func Menu(perso *character.Personnage, inv inventory.Inventory) {
	var choix int
	game.ClearScreen()
	for {
		fmt.Println("\n---------- MENU PRINCIPAL ----------")
		fmt.Println("1 - Aller au combat")
		fmt.Println("2 - Accéder à l'inventaire")
		fmt.Println("3 - Aller au O'Tacos (régénération de vie)")
		fmt.Println("4 - Accéder au Telegram (marché)")
		fmt.Println("5 - Aller zoner dans le quartier")
		fmt.Println("6 - Jouez au morpion avec un ami")
		fmt.Println("0 - Quitter")
		fmt.Println("  - Vos points de vie ")
		game.DisplayHealth(perso.NameUser, perso.Hp, perso.Hpmax)
		fmt.Printf("Vous avez %v€\n", perso.Gold)
		fmt.Println("------------------------------------")
		for {
			fmt.Printf("Votre choix: ")
			fmt.Scan(&choix)
			if choix >= 0 && choix <= 6 {
				break
			} else if choix != 0 && choix != 1 && choix != 2 && choix != 3 && choix != 4 && choix != 5 && choix != 6 {
				fmt.Println("Choix invalide, veuillez entrer une valeur valide.")
			}
		}
		switch choix {
		case 1:
			var choix2 int
			game.ClearScreen()
			fmt.Println("1 - Combat contre le vendeur")
			fmt.Println("2 - Combat contre le Go Fasteur")
			fmt.Println("3 - Combat contre l'homme de main")
			fmt.Println("4 - Combat contre le gérant")
			fmt.Println("5 - Combat contre le caïd")
			fmt.Println("0 - Revenir au menu")
			for {
				fmt.Printf("Votre choix: ")
				fmt.Scan(&choix2)
				if choix2 >= 0 && choix2 < 6 {
					break
				} else if choix2 != 0 && choix2 != 1 && choix2 != 2 && choix2 != 3 && choix2 != 4 && choix2 != 5 {
					fmt.Println("Choix invalide, veuillez entrer une valeur valide.")
				}
			}
			switch choix2 {
			case 1:
				if perso.CombatCounteur >= 1 {
					fmt.Println("Vous vous préparez à entrer dans un combat...")
					time.Sleep(3 * time.Second)
					fight.SecondFight(perso, &inv)
				} else {
					fmt.Println("Veuillez combattre le guetteur d'abord")
				}
			case 2:
				if perso.CombatCounteur >= 2 {
					fmt.Println("Vous vous préparez à entrer dans un combat...")
					time.Sleep(3 * time.Second)
					fight.ThirdFight(perso, &inv)
				} else {
					fmt.Println("Veuillez combattre le vendeur d'abord")
				}
			case 3:
				if perso.CombatCounteur >= 3 {
					fmt.Println("Vous vous préparez à entrer dans un combat...")
					time.Sleep(3 * time.Second)
					fight.FourthFight(perso, &inv)
				} else {
					fmt.Println("Veuillez combattre le GoFasteur d'abord")
				}
			case 4:
				if perso.CombatCounteur >= 4 {
					fmt.Println("Vous vous préparez à entrer dans un combat...")
					time.Sleep(3 * time.Second)
					fight.FifthFight(perso, &inv)
				} else {
					fmt.Println("Veuillez combattre l'Homme de main d'abord")
				}
			case 5:
				if perso.CombatCounteur >= 5 {
					fmt.Println("Vous vous préparez à entrer dans un combat...")
					time.Sleep(3 * time.Second)
					fight.SixthFight(perso, &inv)
				} else {
					fmt.Println("Veuillez combattre le Gérant d'abord")
				}
			case 0:
				break
			}
		case 2:
			game.ClearScreen()
			afficherInventaire(&inv)
		case 3:
			game.ClearScreen()
			oTacos(perso, &inv)
		case 4:
			game.ClearScreen()
			Telegram(perso, &inv)
		case 5:
			go music.PlayMusic("music/secondary_music.mp3")
			game.ClearScreen()
			fmt.Println("Vous zonez dans le quartier !")
			ascii.Jul()
			time.Sleep(15 * time.Second)
			fmt.Println("Vous avez gagner 25€")
			perso.Gold += 25
		case 6:
			game.ClearScreen()
			jeutel.PlayMorpion()
		case 0:
			fmt.Println("Merci et à bientôt !")
			time.Sleep(4 * time.Second)
			os.Exit(0)
		}
	}
}

func afficherInventaire(inv *inventory.Inventory) {
	cocktail := object.ObjectStats{Name: "Cocktail Molotov", Type: "Utilitaire", Damage: 150}
	fmt.Println("\n--------- INVENTAIRE ---------")
	if len(inv.SacocheCp) > 0 {
		for i, obj := range inv.SacocheCp {
			fmt.Printf("%d. %s (%s)\n", i+1, obj.Name, obj.Type)
		}
	} else {
		fmt.Println("Votre inventaire est vide.")
	}
	fmt.Println("\n----- INVENTAIRE DE CRAFT -----")
	if len(inv.CraftInventory) > 0 {
		for i, obj := range inv.CraftInventory {
			fmt.Printf("%d. %s (%s)\n", i+1, obj.Name, obj.Type)
		}
	} else {
		fmt.Println("Votre inventaire d'objet à craft est vide.")
	}
	time.Sleep(2 * time.Second)

	var choix int
	for {
		fmt.Printf("\nTapez 1 pour supprimer un objet ou tapez 2 pour crafter un objet, sinon tapez 0. \n\n")
		fmt.Scan(&choix)
		if choix == 1 || choix == 0 || choix == 2 {
			break
		}
		fmt.Println("Choix invalide, veuillez entrer une valeur valide.")
	}

	switch choix {
	case 1:
		var Nomobjet string
		fmt.Println("Tapez le nom de l'objet à supprimer")
		fmt.Scan(&Nomobjet)
		if Nomobjet == "Briquet" || Nomobjet == "Tissu" || Nomobjet == "Bouteille d'alcool en verre" {
			fmt.Println("Vous ne pouvez pas supprimer les objets de craft")
			return
		}
		// Recherche de l'objet dans l'inventaire
		found := false
		for index, obj := range inv.SacocheCp {
			if obj.Name == Nomobjet {
				// Suppression de l'objet à l'index donné
				inv.SacocheCp = append(inv.SacocheCp[:index], inv.SacocheCp[index+1:]...)
				fmt.Printf("Objet '%s' supprimé avec succès.\n", Nomobjet)
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Objet '%s' non trouvé dans l'inventaire.\n", Nomobjet)
		}
	case 2:
		// Compte des objets de craft dans l'inventaire
		if len(inv.CraftInventory) >= 3 {
			// Si assez d'objets de craft, création du cocktail
			inv.CraftInventory = inv.CraftInventory[3:]
			inv.AddObject(cocktail)
			fmt.Println("Cocktail crafté avec succès !")
		} else {
			fmt.Println("Vous n'avez pas tous les objets requis pour le craft.")
		}
		time.Sleep(3 * time.Second)
		return
	case 0:
		return
	}
}

func oTacos(perso *character.Personnage, inv *inventory.Inventory) {
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
	Menu(perso, *inv)
}

func Telegram(perso *character.Personnage, inv *inventory.Inventory) {
	var achat int
	afficherMarché()

	fmt.Printf("\nVous avez %d€.\n", perso.Gold)

	for {
		fmt.Println("Entrez le numéro de l'article que vous souhaitez acheter, ou 0 pour quitter:")
		fmt.Scan(&achat)
		if achat >= 0 && achat <= 10 {
			break
		}
		fmt.Println("Choix invalide, veuillez entrer une valeur valide.")
	}

	switch achat {
	case 1:
		achatObjet(perso, inv, object.ObjectStats{Name: "Lacrymogène", Type: "Arme", Damage: 15}, 30)
	case 2:
		achatObjet(perso, inv, object.ObjectStats{Name: "Matraque", Type: "Arme", Damage: 80}, 100)
	case 3:
		achatObjet(perso, inv, object.ObjectStats{Name: "Mortier", Type: "Arme", Damage: 200}, 150)
	case 4:
		achatObjet(perso, inv, object.ObjectStats{Name: "Taser", Type: "Arme", Damage: 100}, 250)
	case 5:
		achatObjet(perso, inv, object.ObjectStats{Name: "Ricard", Type: "Soin", Damage: 30}, 5)
	case 6:
		achatObjet(perso, inv, object.ObjectStats{Name: "Flash", Type: "Soin", Damage: 75}, 20)
	case 7:
		achatObjet(perso, inv, object.ObjectStats{Name: "Redbull", Type: "Utilitaire", Damage: 20}, 10)
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
		Menu(perso, *inv)
	}
}

func afficherMarché() {
	fmt.Println("\n--------- TELEGRAM MARCHÉ ---------")
	fmt.Println("1 - LACRIMOGENE à 30€ : -15 point de vie à l'adversaire pendant 3 tours")
	fmt.Println("2 - MATRAQUE à 100€ : Inflige 80 dégâts")
	fmt.Println("3 - MORTIER à 150€ : Inflige 200 dégâts")
	fmt.Println("4 - TASER à 250€ : Inflige 100 dégâts et immobilise 1 tour")
	fmt.Println("5 - RICARD à 5€ : Récupère 30 PV")
	fmt.Println("6 - FLASH à 20€ : Récupère 75 PV")
	fmt.Println("7 - REDBULL à 10€ : Augmente dégâts de 20% (3 attaques)")
	fmt.Println("8 - ENSEMBLE NIKE TECH à 200€ : Augmente PV max de 20")
	fmt.Println("9 - CASQUE ARAI à 500€ : Augmente PV max de 50")
	fmt.Println("10 - SACOCHE LV à 300€ : Augmente la capacité d'inventaire de 5 emplacements")
	fmt.Println("0 - Quitter le Telegram")
	fmt.Println("----------------------------------")
}

func achatObjet(perso *character.Personnage, inv *inventory.Inventory, objet object.ObjectStats, prix int) {
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
