package game

import (
	"fmt"
	"math/rand"
	"projet-red/character"
	"time"
)

func ChoixPersonnage() character.Personnage {
	time.Sleep(3 * time.Second)
	fmt.Println("\n\nDARON : 55 ans, fumeur, passe son temps au PMU, supporte l'OM plus que sa femme. 100 kg au compteur, se bat avec une ceinture. \n(Beaucoup de points de vie, dégâts faibles) TAPEZ 1\n\nDARONNE : 44 ans, femme de ménage, regarde Les Marseillais et TPMP tous les soirs, brune, 60 kg, lance des claquettes pour attaquer. \n(Dégâts élevés, peu de points de vie) TAPEZ 2\n\nTONTON : Tonton cool, il aime les femmes, la plage, les paris sportifs et la pétanque par-dessus tout, gagne sa vie au black et avec l'achat-revente de T-MAX volés,\nrêve d'une carrière de bouliste et ne se bat QU'AVEC SES BOULES UNIQUEMENT. Il est mauvais, mais avec un peu de chance, il peut infliger des dégâts énormes. \n(Points de vie moyens, dégâts aléatoires) TAPEZ 3")
	var choix_perso int
	fmt.Scan(&choix_perso)

	for choix_perso != 1 && choix_perso != 2 && choix_perso != 3 {
		fmt.Printf("Entrez seulement 1, 2 ou 3. \n\n")
		fmt.Scan(&choix_perso)
	}
	var nameuser string
	var perso character.Personnage
	switch choix_perso {
	case 1:
		fmt.Printf("\nChoisisez un pseudo pour votre personnage \n\n")
		fmt.Scan(&nameuser)
		for len(nameuser) > 12 {
			fmt.Printf("\nEntrez maximum 12 caractère !\n \n")
			fmt.Scan(&nameuser)
		}
		perso = character.Personnage{Name: "Daron", NameUser: nameuser, Hp: 120, Gold: 10, Damage: 20, Hpmax: 120, CombatCounteur: 0}
		fmt.Printf("Voici les stats du %v %v \nPoints de vie : %v\nArgent : %v", perso.Name, perso.NameUser, perso.Hp, perso.Gold)
	case 2:
		fmt.Printf("\nChoisisez un pseudo pour votre personnage \n\n")
		fmt.Scan(&nameuser)
		for len(nameuser) > 12 {
			fmt.Printf("\nEntrez maximum 12 caractère !\n \n")
			fmt.Scan(&nameuser)
		}
		perso = character.Personnage{Name: "Daronne", NameUser: nameuser, Hp: 80, Gold: 10, Damage: 35, Hpmax: 80, CombatCounteur: 0}
		fmt.Printf("Voici les stats de la %v %v \nPoints de vie : %v\nArgent : %v", perso.Name, perso.NameUser, perso.Hp, perso.Gold)
	case 3:
		fmt.Printf("\nChoisisez un pseudo pour votre personnage \n \n")
		fmt.Scan(&nameuser)
		for len(nameuser) > 12 {
			fmt.Printf("\nEntrez maximum 12 caractère !\n \n")
			fmt.Scan(&nameuser)
		}
		rand.Seed(time.Now().UnixNano())
		argentrandom := rand.Intn(46) + 5
		perso = character.Personnage{Name: "Tonton", NameUser: nameuser, Hp: 5000, Gold: argentrandom + 5000, Damage: character.DegatTonton(), Hpmax: 100, CombatCounteur: 5}
		fmt.Printf("\nVoici les stats du %v %v \nPoints de vie : %v\nArgent : %v\n", perso.Name, perso.NameUser, perso.Hp, perso.Gold)
	default:
	}
	return perso

}

func DisplayHealth(entityName string, hp int, hpMax int) {
	healthBar := "["
	barLength := 10                        // Longueur totale de la barre
	filledBars := (hp * barLength) / hpMax // Calcul du nombre de barres remplies en fonction des HP actuels et max

	// Ajouter les barres remplies
	for i := 0; i < filledBars; i++ {
		healthBar += "|"
	}
	// Ajouter les barres vides
	for i := filledBars; i < barLength; i++ {
		healthBar += " "
	}
	healthBar += "]"

	// Afficher le nom de l'entité, la barre de santé et les HP actuels/maximum
	fmt.Printf("%s : %s (%d/%d)\n", entityName, healthBar, hp, hpMax)
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
