package fight

import (
	"fmt"
	"projet-red/character"
	"projet-red/game"
	"projet-red/inventory"
	"projet-red/story"
	"time"
)

func SixthFight(perso *character.Personnage, inv inventory.Inventory) inventory.Inventory {
	Caid := character.Enemy{Name: "Le Caid", Hp: 1000, Damage: 35}
	game.ClearScreen()
	story.AfterGerant()
	fmt.Println("\nVous entrez dans un combat avec le Caid !")
	fmt.Println(`
	   O                         O
	  /|\                       /|\
	  / \                       / \`)
	time.Sleep(2 * time.Second)
	for Caid.Hp > 0 && perso.Hp > 0 {
		game.ClearScreen()
		if perso.Name == "Tonton" {
			inv.SacocheCp[0].Damage = character.DegatTonton()
		}
		fmt.Println("\n--- Combat ---")
		fmt.Println("\n--- Combat ---")
		fmt.Println(`
 O                          O
/|\                      //|||\\
/ \                       // \\  `)
		game.DisplayHealth(perso.NameUser, perso.Hp, perso.Hpmax)
		game.DisplayHealth(Caid.Name, Caid.Hp, 1000)
		inv.AfficherInventaireEnCombat()

		// Sélection de l'action par l'utilisateur
		attack := chooseActionCaid(len(inv.SacocheCp), inv)

		// Appliquer l'action choisie
		handleActionCaid(attack, &Caid, perso, &inv)

		// Vérifier si l'un des deux personnages est mort
		if Caid.Hp <= 0 {
			fmt.Println("\n Tu as vaincu le Caïd ! Son corps s’effondre sur le sol tandis que tu t’approches du coffre-fort. À l’intérieur, tu trouves des liasses de billets, des téléphones cryptés, et des preuves qui te permettent de démanteler tout le réseau. Tu es désormais le maître du quartier !")
			perso.Gold += 5000
			if perso.CombatCounteur < 6 {
				perso.CombatCounteur = 6
			}
			return inv
		} else if perso.Hp <= 0 {
			break
		}

		// Riposte du Homme de main
		enemyRetaliationCaid(&Caid, perso)
	}
	fmt.Println("\n Le Caïd t'a battu. Tu te réveilles à l’hôpital, gravement blessé. Le quartier est à nouveau sous contrôle du réseau. Tu devras revenir plus fort pour avoir une chance de les détrôner")
	perso.Hp = perso.Hpmax / 2
	perso.Gold /= 2
	time.Sleep(5 * time.Second)
	return inv
}

// Choisir une action valide
func chooseActionCaid(max int, inv inventory.Inventory) int {
	var attack int
	// Parcours des objets dans l'inventaire et affiche les options disponibles.
	for i := 0; i < len(inv.SacocheCp); i++ {
		// Si l'objet n'est pas "Utilitaire", il est soit une arme soit un soin.
		if inv.SacocheCp[i].Type != "Utilitaire" {
			if inv.SacocheCp[i].Type == "Soin" {
				// Affiche les options pour les objets de soin.
				fmt.Printf("\nPour utiliser le %v et soigner %v pv. Appuyez sur %v\n", inv.SacocheCp[i].Name, inv.SacocheCp[i].Damage, i)
			}
			if inv.SacocheCp[i].Type == "Arme" {
				// Affiche les options pour les armes.
				if inv.SacocheCp[i].Name == "Boule" {
					fmt.Printf("\nPour utiliser la %v et infliger des dégâts aléatoire. Appuyez sur %v\n", inv.SacocheCp[i].Name, i)
				} else {
					fmt.Printf("\nPour utiliser le %v et infliger %v dégâts. Appuyez sur %v\n", inv.SacocheCp[i].Name, inv.SacocheCp[i].Damage, i)
				}
			}
		}
		if inv.SacocheCp[i].Name == "Cocktail Molotov" {
			fmt.Printf("\nPour utiliser le %v et infliger %v dégâts à tout les adversaires (1 utilisation possible), Appuyez sur %v\n", inv.SacocheCp[i].Name, inv.SacocheCp[i].Damage, i)
		}
		if inv.SacocheCp[i].Name == "Redbull" {
			fmt.Printf("\nPour utilisez le %v et augmenter vos dégâts %v%%. Appuyez sur %v\n", inv.SacocheCp[i].Name, inv.SacocheCp[i].Damage, i)
		}
	}
	// Attente de l'entrée utilisateur pour choisir une action valide.
	for {
		fmt.Print("\nEntrez votre choix : ")
		fmt.Scan(&attack)
		if attack >= 0 && attack < max {
			return attack
		}
		fmt.Println("Entrez une option valide.")
	}
}

// Gérer l'action choisie par le joueur (attaque ou soin)
func handleActionCaid(attack int, enemy *character.Enemy, perso *character.Personnage, inv *inventory.Inventory) {
	item := inv.SacocheCp[attack]
	if item.Name == "Redbull" {
		// Gestion du Redbull qui augmente les dégâts pendant 3 tours
		fmt.Println("Vous buvez un Redbull, vos dégâts sont augmentés pendant 3 tours !")
		perso.BoostDamageTurns = 4
		inv.RemoveObject(item) // Supprimer l'objet après usage
		return
	}
	if item.Name == "Cocktail Molotov" {
		fmt.Println("Vous infligez des dégâts seulement au Caid")
		enemy.Hp -= 150
		inv.RemoveObject(item)
		return
	}

	if item.Type == "Arme" {
		// Gestion de la Lacrymogène
		if item.Name == "Lacrymogène" {
			enemy.LacrymogèneActive = true
			enemy.LacrymogèneTurns = 3
			fmt.Printf("Vous lancez une lacrymogène sur %s. Il subira 15 points de dégâts par tour pendant 3 tours !\n", enemy.Name)
			inv.RemoveObject(item) // Suppression de l'objet après utilisation
			return
		}

		if item.Name == "Taser" {
			enemy.Immobilized = true
			enemy.ImmobilizedTurns = 2
			fmt.Printf("Vous utilisez le Taser sur %s. Il est immobilisé pour 1 tour !\n", enemy.Name)
			inv.RemoveObject(item)
			return
		}
		damage := item.Damage
		damageFloat := float32(damage) * 1.2
		if perso.BoostDamageTurns > 0 {
			damage = int(damageFloat)
		}
		enemy.Hp -= damage
		if enemy.Hp < 0 {
			enemy.Hp = 0
		}
		fmt.Printf("Vous infligez %d points de dégât à %s.\n", damage, enemy.Name)

	} else if item.Type == "Soin" {
		inv.RemoveObject(item)
		heal := item.Damage
		perso.Hp += heal
		fmt.Printf("Vous vous soignez de %d pv.\n", heal)
	}
	time.Sleep(2 * time.Second)
}

// Riposte de l'ennemi
func enemyRetaliationCaid(enemy *character.Enemy, perso *character.Personnage) {
	time.Sleep(1 * time.Second)

	// Réduire le nombre de tours pour le boost de dégâts du joueur (Redbull)
	if perso.BoostDamageTurns > 0 {
		perso.BoostDamageTurns--
		if perso.BoostDamageTurns == 0 {
			fmt.Println("L'effet de Redbull est terminé, vos dégâts reviennent à la normale.")
		}
	}

	if enemy.Immobilized {
		enemy.ImmobilizedTurns--
		if enemy.ImmobilizedTurns == 1 {
			enemy.Damage = 0
			fmt.Printf("%s est immobilisé et ne peut pas attaquer ce tour.\n", enemy.Name)
		}
		if enemy.ImmobilizedTurns <= 0 {
			enemy.Damage = 35
			enemy.Immobilized = false
			fmt.Printf("%s n'est plus immobilisé.\n", enemy.Name)
		}
	}

	if enemy.LacrymogèneActive {
		enemy.Hp -= 15
		if enemy.Hp < 0 {
			enemy.Hp = 0
		}
		fmt.Printf("%s subit 15 points de dégâts à cause de la lacrymogène.\n", enemy.Name)
		enemy.LacrymogèneTurns--
		if enemy.LacrymogèneTurns <= 0 {
			enemy.LacrymogèneActive = false
			fmt.Printf("L'effet de la lacrymogène sur %s a pris fin.\n", enemy.Name)
		}
	}

	time.Sleep(1 * time.Second)
	perso.Hp -= enemy.Damage
	fmt.Printf("Le Caid vous inflige %d points de dégât.\n", enemy.Damage)
	time.Sleep(2 * time.Second)
}
