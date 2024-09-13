package game

import (
	"fmt"
	"projet-red/character"
	"time"
)

func Testcombat() {
	attack := 0
	guetteur := character.Enemy{"Guetteur", 100, 10}
	perso := character.Personnage{"Daron", "oui", 120, 15, 20}
	fmt.Println("\nVous rencontrez un guetteur et décidez de lui arracher la gueule")
	fmt.Println(`
   O                         O
  /|\                       /|\
  / \                       / \
`)
	for guetteur.Hp > 0 && perso.Hp > 0 {

		fmt.Println("\n", perso.Name, "a", perso.Hp, "point de vie et ", guetteur.Name, "a", guetteur.Hp, "point de vie")
		fmt.Println("\nAppuyez sur 1 pour lui peter la gueule")
		fmt.Scan(&attack)
		for attack != 1 {
			fmt.Println("Entrez une valeur valide\n")
			fmt.Scan(&attack)
		}
		switch attack {

		case 1:
			if guetteur.Hp > 21 {
				guetteur.Hp -= perso.Damage
				fmt.Println("Vous attaquez le guetteur et vous lui infligez 10 points de dégât")
				time.Sleep(2 * time.Second)
				perso.Hp -= guetteur.Damage
				fmt.Println("\nLe guetteur vous attaque en retour et vous enlève 10 points de vie")
			} else {
				fmt.Print("\nLe guetteur sort un couteau et vous inflige 50 points de dégât")
				perso.Hp -= 50
				guetteur.Hp -= perso.Damage
			}
		}
	}
	if guetteur.Hp <= 0 {
		time.Sleep(2 * time.Second)
		fmt.Println("\n", perso.Name, "a", perso.Hp, "point de vie et ", guetteur.Name, "a", guetteur.Hp, "point de vie")
		time.Sleep(2 * time.Second)
		fmt.Println("\nVous avez gagnez vôtre combat")
	} else if perso.Hp <= 0 {
		time.Sleep(2 * time.Second)
		fmt.Println("\n", perso.Name, "a", perso.Hp, "point de vie et ", guetteur.Name, "a", guetteur.Hp, "point de vie")
		fmt.Println("\nTu t'es fait arraché t'es nul")
	}
}
