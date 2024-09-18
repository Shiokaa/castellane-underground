package fight

import (
	"fmt"
	"projet-red/character"
	"time"
)

func Secondfight(perso character.Personnage) {
	attack := 0
	vendeur := character.Enemy{Name: "vandeur", Hp: 200, Damage: 20}
	fmt.Println("\nTrès bien ! La discussion prend fin et tu décides de te rendre à l’entrée du quartier afin de rencontrer le guetteur pour récolter des informations sur ses supérieurs, mais celui-ci n’est pas ouvert au débat. Le combat commence !")
	fmt.Println(`
	   O                         O
	  /|\                       /|\
	  / \                       / \`)
	for vendeur.Hp > 0 && perso.Hp > 0 {

		fmt.Println("\n", perso.NameUser, "a", perso.Hp, "point de vie et ", vendeur.Name, "a", vendeur.Hp, "point de vie")
		fmt.Println("\nAppuyez sur 1 pour lui peter la gueule")
		fmt.Scan(&attack)
		for attack != 1 {
			fmt.Printf("Entrez une valeur valide\n\n")
			fmt.Scan(&attack)
		}

		switch attack {
		case 1:
			damage := perso.Damage
			vendeur.Hp -= damage
			fmt.Printf("Vous infligez %d points de dégât.\n", damage)
			time.Sleep(2 * time.Second)
			fmt.Println("Le vendeur riposte !")
			time.Sleep(1 * time.Second)
			perso.Hp -= vendeur.Damage
			fmt.Printf("Le vendeur vous inflige %d points de dégât.\n", vendeur.Damage)
			time.Sleep(2 * time.Second)

			if perso.Hp <= 0 {
				fmt.Println("\nVous êtes tombé au combat...")
			} else if vendeur.Hp <= 0 {
				fmt.Println("\nVous avez vaincu le vendeur ! Vous trouvez du tissu et  100 euros !!")
				/* inv.AddObject(tissu) */
				perso.Gold += 100
			}
		}
	}
}


inv := inventory.Inventory{Limite: 5}
briquet := object.ObjectStats{Name: "Briquet", Type: "Utilitaire", Damage: 10}
sandwitch := object.ObjectStats{Name: "Sandwitch", Type: "Soin", Damage: 10}
ricard := object.ObjectStats{Name: "Ricard", Type: "Soin", Damage: 10}

attack := 0
guetteur := character.Enemy{Name: "Guetteur", Hp: 100, Damage: 10}

fmt.Println("\nVous entrez dans un combat avec un Guetteur !")
fmt.Println(`
O                         O
/|\                       /|\
/ \                       / \`)
time.Sleep(2 * time.Second)

for guetteur.Hp > 0 && perso.Hp > 0 {
	fmt.Println("\n--- Combat ---")

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
		damage := perso.Damage
		vendeur.Hp -= damage
		fmt.Printf("Vous infligez %d points de dégât.\n", damage)
		time.Sleep(2 * time.Second)

		if vendeur.Hp > 19 {
			fmt.Println("Le vendeur riposte !")
			time.Sleep(1 * time.Second)
			perso.Hp -= vendeur.Damage
			fmt.Printf("Le vendeur vous inflige %d points de dégât.\n", guetteur.Damage)
			time.Sleep(3 * time.Second)
			game.ClearScreen()
		} 
		time.Sleep(2 * time.Second)

		if perso.Hp <= 0 {
			fmt.Println("\nVous êtes tombé au combat...")
		} else if vendeur.Hp <= 0 {
			fmt.Println("\nVous avez vaincu le vendeur !")
			inv.AddObject(briquet)
			inv.AddObject(sandwitch)
			inv.AddObject(ricard)
			break
		}
	}

}
return inv
}