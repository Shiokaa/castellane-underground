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
			fmt.Printf("Le vendeur vous inflige %d points de dégât.\n", vendeur.eDamage)
			time.Sleep(2 * time.Second)

			if perso.Hp <= 0 {
				fmt.Println("\nVous êtes tombé au combat...")
			} else if vendeur.Hp <= 0 {
				fmt.Println("\nVous avez vaincu le vendeur ! Vous trouvez du tissu et  100 euros !!")
				inv.AddObject(tissu)
				perso.Gold += 100
			}
		}
	}
}

