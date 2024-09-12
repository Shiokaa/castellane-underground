package game

import (
	"fmt"
	"projet-red/character"
)

func ChoixPersonnage() character.Personnage {
	fmt.Println("\nBienvenue dans le menu du choix des personnages. Qui allez vous choisir pour renverser les réseaux du quartier ?! \n\nTapez 1 pour selectionner le daron. \nTapez 2 pour selectionner la daronne. \nTapez 3 pour selectionner le tonton.")
	var choix_perso int
	fmt.Scan(&choix_perso)
	for choix_perso != 1 && choix_perso != 2 && choix_perso != 3 {
		fmt.Println("Entrez seulement 1,2 ou 3. \n")
		fmt.Scan(&choix_perso)
	}
	if choix_perso == 1 {
		perso := character.Personnage{"Daron", 120, 20}
		fmt.Printf("Voici les stats du %v \nPoints de vie : %v", perso.Name, perso.Hp)
		return perso
	}
	if choix_perso == 2 {
		perso := character.Personnage{"Daronne", 80, 20}
		fmt.Printf("Voici les stats de la %v \nPoints de vie : %v", perso.Name, perso.Hp)
		return perso
	}
	if choix_perso == 3 {
		perso := character.Personnage{"Tonton", 80, 20}
		fmt.Printf("Voici les stats du %v \nPoints de vie : %v", perso.Name, perso.Hp)
		return perso
	}
	return character.Personnage{"Tonton", 80, 20}
}
