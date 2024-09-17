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
			vendeur.Hp -= perso.Damage
			fmt.Printf("Vous attaquez le vendeur et vous lui infligez %v points de dégât", perso.Damage)
			time.Sleep(2 * time.Second)
			perso.Hp -= vendeur.Damage
			fmt.Printf("\nLe vendeur vous attaque en retour et vous enlève %v points de vie", vendeur.Damage)
		}

		if vendeur.Hp <= 0 {
			time.Sleep(2 * time.Second)
			fmt.Println("\n", perso.Name, "a", perso.Hp, "point de vie et ", vendeur.Name, "a", vendeur.Hp, "point de vie")
			time.Sleep(2 * time.Second)
			fmt.Println("\nVous avez gagnez vôtre combat? Vous dépouyer le vendeur et trouvez 100 euros ! ")
			perso.Gold += 100
			fmt.Printf("Vous possedez désormais %v euros", perso.Gold)
		} else if perso.Hp <= 0 {
			time.Sleep(2 * time.Second)
			fmt.Println("\n", perso.Name, "a", perso.Hp, "point de vie et ", vendeur.Name, "a", vendeur.Hp, "point de vie")
			fmt.Println("\nTu t'es fait arraché t'es nul")
		}
	}
}
