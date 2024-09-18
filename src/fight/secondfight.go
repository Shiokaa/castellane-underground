package fight

import (
	"fmt"
	"projet-red/character"
	"projet-red/game"
	"time"
)

func Secondfight(perso character.Personnage) {
	attack := 0
	Vendeur := character.Enemy{Name: "Vendeur", Hp: 200, Damage: 20}
	fmt.Println("\n Vous entrez dans un combat avec le vendeur !")
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

		fmt.Println("\n", perso.NameUser, "a", perso.Hp, "point de vie et ", Vendeur.Name, "a", Vendeur.Hp, "point de vie")
		fmt.Println("\nAppuyez sur 1 pour lui peter la gueule")
		fmt.Scan(&attack)
		for attack != 1 {
			fmt.Printf("Entrez une valeur valide\n\n")
			fmt.Scan(&attack)
		}

		switch attack {
		case 0:
			damage := perso.Damage
			Vendeur.Hp -= damage
			fmt.Printf("Vous infligez %d points de dégât.\n", damage)
			time.Sleep(2 * time.Second)
			fmt.Println("Le vendeur riposte !")
			time.Sleep(1 * time.Second)
			perso.Hp -= Vendeur.Damage
			fmt.Printf("Le vendeur vous inflige %d points de dégât.\n", Vendeur.Damage)
			time.Sleep(2 * time.Second)

			if perso.Hp <= 0 {
				fmt.Println("\nVous êtes tombé au combat...")
			} else if Vendeur.Hp <= 0 {
				fmt.Println("\nVous avez vaincu le vendeur ! Vous trouvez du tissu et  100 euros !!")
				/* inv.AddObject(tissu) */
				perso.Gold += 100
			}
		}
	}
}
