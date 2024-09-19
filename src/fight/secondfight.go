package fight

import (
	"fmt"
	"projet-red/character"
	"projet-red/game"
	"projet-red/inventory"
	"projet-red/object"
	"time"
)

func SecondFight(perso *character.Personnage, inv inventory.Inventory) inventory.Inventory {
	Vendeur := character.Enemy{Name: "Vendeur", Hp: 300, Damage: 20}
	tissu := object.ObjectStats{Name: "tissu", Type: "Utilitaire", Damage: 0}

	fmt.Println("\nVous entrez dans un combat avec un Vendeur !")
	fmt.Println(`
	   O                         O
	  /|\                       /|\
	  / \                       / \`)
	time.Sleep(2 * time.Second)

	for Vendeur.Hp > 0 && perso.Hp > 0 {
		fmt.Println("\n--- Combat ---")
		game.DisplayHealth(perso.NameUser, perso.Hp, perso.Hpmax)
		game.DisplayHealth(Vendeur.Name, Vendeur.Hp, 100)
		inv.AfficherInventaireEnCombat()

		// Sélection de l'action par l'utilisateur
		attack := chooseAction(len(inv.SacocheCp), inv)

		// Appliquer l'action choisie
		handleAction(attack, &Vendeur, perso, inv)

		// Vérifier si l'un des deux personnages est mort
		if Vendeur.Hp <= 0 {
			fmt.Println("\nVous avez vaincu le Vendeur ! Vous trouvez du tissu et 100 euros !!")
			inv.AddObject(tissu)
			perso.Gold += 100
			break
		} else if perso.Hp <= 0 {
			fmt.Println("\nVous êtes tombé au combat...")
			break
		}

		// Riposte du Vendeur
		enemyRetaliation(&Vendeur, perso)
	}

	return inv
}

// Choisir une action valide
func chooseAction(max int, inv inventory.Inventory) int {
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
func handleAction(attack int, enemy *character.Enemy, perso *character.Personnage, inv inventory.Inventory) {
	item := inv.SacocheCp[attack]
	if item.Type == "Arme" {
		damage := item.Damage
		enemy.Hp -= damage
		fmt.Printf("Vous infligez %d points de dégât à %s.\n", damage, enemy.Name)
	} else if item.Type == "Soin" {
		heal := item.Damage
		perso.Hp += heal
		fmt.Printf("Vous vous soignez de %d pv.\n", heal)
	}
	time.Sleep(2 * time.Second)
}

// Riposte de l'ennemi
func enemyRetaliation(enemy *character.Enemy, perso *character.Personnage) {
	fmt.Println("Le Vendeur riposte !")
	time.Sleep(1 * time.Second)
	perso.Hp -= enemy.Damage
	fmt.Printf("Le Go fasteur vous inflige %d points de dégât.\n", enemy.Damage)
	time.Sleep(2 * time.Second)
}
