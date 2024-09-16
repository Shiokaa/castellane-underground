package fight

import (
	"fmt"
	"projet-red/character"
	"projet-red/game"
	"projet-red/inventory"
	"projet-red/object"
	"time"
)

func Firstfight(perso *character.Personnage) inventory.Inventory {
	inv := inventory.Inventory{Limite: 5}
	briquet := object.ObjectStats{"Briquet", "Utilitaire", 10}
	sandwitch := object.ObjectStats{"Sandwitch", "Utilitaire", 10}
	ricard := object.ObjectStats{"Ricard", "Utilitaire", 10}

	attack := 0
	guetteur := character.Enemy{"Guetteur", 100, 10}

	fmt.Println("\nVous entrez dans un combat avec un Guetteur !")
	fmt.Println(`
   O                         O
  /|\                       /|\
  / \                       / \
`)
	time.Sleep(2 * time.Second)

	for guetteur.Hp > 0 && perso.Hp > 0 {
		fmt.Println("\n--- Combat ---")

		game.DisplayHealth(perso.NameUser, perso.Hp, perso.Hpmax)
		game.DisplayHealth(guetteur.Name, guetteur.Hp, 100)

		fmt.Println("\nQue voulez-vous faire ?")
		fmt.Println("1 - Attaquer")
		fmt.Scan(&attack)

		for attack != 1 && attack != 2 {
			fmt.Println("Entrez une option valide (1)\n")
			fmt.Scan(&attack)
		}

		switch attack {
		case 1:
			damage := perso.Damage
			guetteur.Hp -= damage
			fmt.Printf("Vous infligez %d points de dégât.\n", damage)
			time.Sleep(2 * time.Second)

			if guetteur.Hp > 19 {
				fmt.Println("Le guetteur riposte !")
				time.Sleep(1 * time.Second)
				perso.Hp -= guetteur.Damage
				fmt.Printf("Le guetteur vous inflige %d points de dégât.\n", guetteur.Damage)
			} else {
				criticalDamage := 50
				perso.Hp -= criticalDamage
				fmt.Println("Le guetteur sort un couteau !")
				fmt.Printf("Il vous inflige un coup critique de %d points de dégât.\n", criticalDamage)
			}
			time.Sleep(2 * time.Second)

			if perso.Hp <= 0 {
				fmt.Println("\nVous êtes tombé au combat...")
				break
			} else if guetteur.Hp <= 0 {
				fmt.Println("\nVous avez vaincu le guetteur !")
				inv.AddObject(briquet)
				inv.AddObject(sandwitch)
				inv.AddObject(ricard)
				break
			}
		}

	}
	return inv
}
