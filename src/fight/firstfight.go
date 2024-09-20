package fight

import (
	"fmt"
	"os"
	"projet-red/character"
	"projet-red/game"
	"projet-red/inventory"
	"projet-red/object"
	"time"
)

func Firstfight(perso *character.Personnage) inventory.Inventory {
	inv := inventory.Inventory{Limite: 5}
	briquet := object.ObjectStats{Name: "Briquet", Type: "Utilitaire", Damage: 10}
	sandwitch := object.ObjectStats{Name: "Sandwitch", Type: "Soin", Damage: 10}
	ricard := object.ObjectStats{Name: "Ricard", Type: "Soin", Damage: 10}
	armeBaseDaron := object.ObjectStats{Name: "Ceinture", Type: "Arme", Damage: 20}
	armeBaseDaronne := object.ObjectStats{Name: "Claquette", Type: "Arme", Damage: 35}
	armeBaseTonton := object.ObjectStats{Name: "Boule", Type: "Arme", Damage: character.DegatTonton()}

	attack := 0
	guetteur := character.Enemy{Name: "Guetteur", Hp: 100, Damage: 10}
	game.ClearScreen()

	fmt.Println("\nVous entrez dans un combat avec un Guetteur !")
	fmt.Println(`
   O                         O
  /|\                       /|\
  / \                       / \`)
	time.Sleep(2 * time.Second)

	for guetteur.Hp > 0 && perso.Hp > 0 {

		fmt.Println("\n--- Combat ---")
		fmt.Println(`
		O                         O
	   /|\                       /|\
	   / \                       / \`)

		game.DisplayHealth(perso.NameUser, perso.Hp, perso.Hpmax)
		game.DisplayHealth(guetteur.Name, guetteur.Hp, 100)

		fmt.Println("\nQue voulez-vous faire ?")
		fmt.Println("1 - Attaquer")
		fmt.Scan(&attack)

		for attack != 1 && attack != 2 {
			fmt.Printf("Entrez une option valide (1)\n\n")
			fmt.Scan(&attack)
		}

		switch attack {
		case 1:
			if perso.Name == "Tonton" {
				perso.Damage = character.DegatTonton()
			}
			guetteur.Hp -= perso.Damage
			fmt.Printf("Vous infligez %d points de dégât.\n", perso.Damage)
			time.Sleep(2 * time.Second)

			if guetteur.Hp > 19 {
				fmt.Println("Le guetteur riposte !")
				time.Sleep(1 * time.Second)
				perso.Hp -= guetteur.Damage
				fmt.Printf("Le guetteur vous inflige %d points de dégât.\n", guetteur.Damage)
				time.Sleep(2 * time.Second)
				game.ClearScreen()
			} else {
				criticalDamage := 50
				perso.Hp -= criticalDamage
				fmt.Println("Le guetteur sort un couteau !")
				fmt.Printf("Il vous inflige un coup critique de %d points de dégât.\n", criticalDamage)
				time.Sleep(2 * time.Second)
				game.ClearScreen()
			}
			time.Sleep(1 * time.Second)

			if perso.Hp <= 0 {
				fmt.Println("\nVous êtes mort sur le premier combat c'est honteux")
				time.Sleep(3 * time.Second)
				os.Exit(0)
			} else if guetteur.Hp <= 0 {
				fmt.Println("\nVous avez vaincu le guetteur !")
				if perso.CombatCounteur < 1 {
					perso.CombatCounteur += 1
				}
				perso.Gold += 10
				if perso.Name == "Daron" {
					inv.AddObject(armeBaseDaron)
				} else if perso.Name == "Daronne" {
					inv.AddObject(armeBaseDaronne)
				} else if perso.Name == "Tonton" {
					inv.AddObject(armeBaseTonton)
				}

				inv.AddCraft(briquet)
				inv.AddObject(sandwitch)
				inv.AddObject(ricard)
				break
			}
		}

	}
	return inv
}
