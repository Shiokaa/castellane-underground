package fight

import (
	"fmt"
	"projet-red/character"
	"projet-red/game"
	"projet-red/inventory"
	"time"
)

var Gérant = character.Enemy{Name: "Gérant", Hp: 300, Damage: 25}
var Guetteur = character.Enemy{Name: "Guetteur", Hp: 100, Damage: 10}
var Guetteur2 = character.Enemy{Name: "Guetteur", Hp: 100, Damage: 10}

func FifthFight(perso *character.Personnage, inv inventory.Inventory) inventory.Inventory {

	fmt.Println("\nVous entrez dans un combat avec un Gérant !")
	fmt.Println(`
	   O                         O
	  /|\                       /|\
	  / \                       / \`)
	time.Sleep(2 * time.Second)

	for Gérant.Hp > 0 && perso.Hp > 0 {
		fmt.Println("\n--- Combat ---")
		game.DisplayHealth(perso.NameUser, perso.Hp, perso.Hpmax)
		game.DisplayHealth(Gérant.Name, Gérant.Hp, 300)
		game.DisplayHealth(Guetteur.Name, Guetteur.Hp, 100)
		game.DisplayHealth(Guetteur2.Name, Guetteur.Hp, 100)
		// Sélection de l'action par l'utilisateur
		attack := chooseAction2(len(inv.SacocheCp), inv)

		// Appliquer l'action choisie
		handleAction2(attack, &Gérant, perso, inv)

		// Vérifier si l'un des deux personnages est mort
		if Gérant.Hp <= 0 {
			fmt.Println("\nVous avez vaincu le Gérant ! Vous trouvez 500 euros !!")
			if perso.CombatCounteur < 4 {
				perso.CombatCounteur = 4
			}
			perso.Gold += 500
			break
		} else if perso.Hp <= 0 {
			fmt.Println("\nVous êtes tombé au combat...")
			break
		}

		// Riposte du Gérant
		enemyRetaliation2(&Gérant, perso)
	}

	return inv
}

func chooseAction2(max int, inv inventory.Inventory) int {
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

func handleAction2(attack int, enemy *character.Enemy, perso *character.Personnage, inv inventory.Inventory) {
	var attack2 int
	item := inv.SacocheCp[attack]
	if item.Type == "Arme" {
		fmt.Print("\nQui souhaitez vous attaquer ?")
		fmt.Print("1- Le Gérant")
		fmt.Print("2- Le Guetteur 1")
		fmt.Print("3- Le Guetteur 2")
		fmt.Scan(&attack2)
		switch attack2 {
		case 1:
			damage := item.Damage
			enemy.Hp -= damage
			fmt.Printf("Vous infligez %d points de dégât à %s.\n", damage, enemy.Name)
		case 2:
			damage := item.Damage
			Guetteur.Hp -= damage
			fmt.Printf("Vous infligez %d points de dégât à %s.\n", damage, Guetteur.Name)

		case 3:
			damage := item.Damage
			Guetteur2.Hp -= damage
			fmt.Printf("Vous infligez %d points de dégât à %s.\n", damage, Guetteur2.Name)
		}
	} else if item.Type == "Soin" {
		heal := item.Damage
		perso.Hp += heal
		fmt.Printf("Vous vous soignez de %d pv.\n", heal)
	}
	time.Sleep(2 * time.Second)
}

// Riposte de l'ennemi
func enemyRetaliation2(enemy *character.Enemy, perso *character.Personnage) {
	fmt.Println("Le Gérant et ses guetteurs ripostent !")
	time.Sleep(1 * time.Second)
	perso.Hp -= enemy.Damage
	fmt.Printf("Le Gérant et ses guetteurs vous inflige %d points de dégât.\n", enemy.Damage)
	time.Sleep(2 * time.Second)
}
