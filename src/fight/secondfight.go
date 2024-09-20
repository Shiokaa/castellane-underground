package fight

import (
	"fmt"
	"projet-red/character"
	"projet-red/game"
	"projet-red/inventory"
	"time"
)

func SecondFight(perso *character.Personnage, inv inventory.Inventory) inventory.Inventory {
	game.ClearScreen()
	Vendeur := character.Enemy{Name: "Vendeur", Hp: 100, Damage: 20}

	fmt.Println("\nTu t’enfonce dans le quartier et fini par trouvé le Vendeur, il refuse de te vendre quoi que se soit et te manque de respect, Le Combat commence !")
	fmt.Println(`
	   O                         O
	  /|\                       /|\
	  / \                       / \`)
	time.Sleep(2 * time.Second)
	inv.AfficherInventaireEnCombat()
	for Vendeur.Hp > 0 && perso.Hp > 0 {
		game.ClearScreen()
		if perso.Name == "Tonton" {
			perso.Damage = character.DegatTonton()
		}
		fmt.Println("\n--- Combat ---")
		game.DisplayHealth(perso.NameUser, perso.Hp, perso.Hpmax)
		game.DisplayHealth(Vendeur.Name, Vendeur.Hp, 100)

		// Sélection de l'action par l'utilisateur
		attack := chooseActionVendeur(len(inv.SacocheCp), inv)

		// Appliquer l'action choisie
		handleActionVendeur(attack, &Vendeur, perso, inv)

		// Vérifier si l'un des deux personnages est mort
		if Vendeur.Hp <= 0 {
			fmt.Println("\nBravo, tu as hagar le revendeur ! , tu le dépouille et en tire 100€ Continue dans le quartier ")
			perso.Gold += 100
			if perso.CombatCounteur < 2 {
				perso.CombatCounteur += 1
			}
			break
		} else if perso.Hp <= 0 {
			fmt.Println("\nTu t’es fait Hagar ! Le vendeur a pris ton argent et t’a envoyé à l’hôpital Nord. Régénère ta vie, puis reviens plus fort !")
			perso.Hp = perso.Hpmax / 2
			perso.Gold /= 2
			time.Sleep(5 * time.Second)
			break
		}

		// Riposte du Vendeur
		enemyRetaliationVendeur(&Vendeur, perso)
	}

	return inv
}

// Choisir une action valide
func chooseActionVendeur(max int, inv inventory.Inventory) int {
	var attack int
	for i := 0; i < len(inv.SacocheCp); i++ {
		if inv.SacocheCp[i].Type != "Utilitaire" {
			if inv.SacocheCp[i].Type == "Soin" {
				fmt.Printf("\nPour utilisez le %v et soigner %v pv. Appuyez sur %v\n", inv.SacocheCp[i].Name, inv.SacocheCp[i].Damage, i)
			}
			if inv.SacocheCp[i].Type == "Arme" {
				if inv.SacocheCp[i].Name == "Boule" {
					fmt.Printf("\nPour utilisez le %v et infliger des degats aléatoire. Appuyez sur %v\n", inv.SacocheCp[i].Name, i)
				} else {
					fmt.Printf("\nPour utilisez le %v et infliger %v deg. Appuyez sur %v\n", inv.SacocheCp[i].Name, inv.SacocheCp[i].Damage, i)
				}
			}
		}
	}
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
func handleActionVendeur(attack int, enemy *character.Enemy, perso *character.Personnage, inv inventory.Inventory) {
	item := inv.SacocheCp[attack]
	if item.Name == "Redbull" {
		// Gestion du Redbull qui augmente les dégâts pendant 3 tours
		fmt.Println("Vous buvez un Redbull, vos dégâts sont augmentés pendant 3 tours !")
		perso.BoostDamageTurns = 4
		inv.RemoveObject(item) // Supprimer l'objet après usage
		return
	}
	if item.Type == "Arme" {
		if item.Name == "Taser" {
			enemy.Immobilized = true
			enemy.ImmobilizedTurns = 2
			fmt.Printf("Vous utilisez le Taser sur %s. Il est immobilisé pour 1 tour !\n", enemy.Name)
			inv.RemoveObject(item)
			return
		}
		// Gestion de la Lacrymogène
		if item.Name == "Lacrymogène" {
			enemy.LacrymogèneActive = true
			enemy.LacrymogèneTurns = 3
			fmt.Printf("Vous lancez une lacrymogène sur %s. Il subira 15 points de dégâts par tour pendant 3 tours !\n", enemy.Name)
			inv.RemoveObject(item) // Suppression de l'objet après utilisation
			return
		}
		damage := item.Damage
		enemy.Hp -= damage
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
func enemyRetaliationVendeur(enemy *character.Enemy, perso *character.Personnage) {
	fmt.Println("Le Vendeur riposte !")
	time.Sleep(1 * time.Second)
	perso.Hp -= enemy.Damage
	fmt.Printf("Le Vendeur vous inflige %d points de dégât.\n", enemy.Damage)
	time.Sleep(2 * time.Second)
}
