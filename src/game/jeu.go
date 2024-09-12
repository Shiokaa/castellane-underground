package game

import (
	"fmt"
	"projet-red/character"
	"time"
)

func HistoireDebut(str string) {
	for _, char := range str {
		fmt.Printf("%c", char)
		time.Sleep(20 * time.Millisecond)
	}
}

func ChoixPersonnage() character.Personnage {
	time.Sleep(3 * time.Second)
	fmt.Println("\n\nBienvenue dans le menu du choix des personnages. Qui allez-vous choisir pour renverser les réseaux du quartier ?! \n\nTapez 1 pour sélectionner le Daron. \nTapez 2 pour sélectionner la Daronne. \nTapez 3 pour sélectionner le Tonton.")

	var choix_perso int
	fmt.Scan(&choix_perso)

	for choix_perso != 1 && choix_perso != 2 && choix_perso != 3 {
		fmt.Println("Entrez seulement 1, 2 ou 3. \n")
		fmt.Scan(&choix_perso)
	}

	var perso character.Personnage
	switch choix_perso {
	case 1:
		perso = character.Personnage{"Daron", 120, 20}
		fmt.Printf("Voici les stats du %v \nPoints de vie : %v\n", perso.Name, perso.Hp)
	case 2:
		perso = character.Personnage{"Daronne", 80, 20}
		fmt.Printf("Voici les stats de la %v \nPoints de vie : %v\n", perso.Name, perso.Hp)
	case 3:
		perso = character.Personnage{"Tonton", 80, 20}
		fmt.Printf("Voici les stats du %v \nPoints de vie : %v\n", perso.Name, perso.Hp)
	default:
		perso = character.Personnage{"Tonton", 80, 20}

	}
	return perso

}
