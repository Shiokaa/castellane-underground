package game

import (
	"fmt"
	"math/rand"
	"projet-red/character"
	"time"
)

<<<<<<< HEAD
func HistoireDebut() {
	fmt.Println("   _____          _       _ _                     _    _           _                                          _ ")
	fmt.Println("  / ____|        | |     | | |                   | |  | |         | |                                        | |")
	fmt.Println(" | |     __ _ ___| |_ ___| | | __ _ _ __   ___   | |  | |_ __   __| | ___ _ __ __ _ _ __ ___  _   _ _ __   __| |")
	fmt.Println(" | |    / _` / __| __/ _ | | |/ _` | '_ \\ / _ \\  | |  | | '_ \\ / _` |/ _ | '__/ _` | '__/ _ \\| | | | '_ \\ / _` |")
	fmt.Println(" | |___| (_| \\__ | ||  __| | | (_| | | | |  __/  | |__| | | | | (_| |  __| | | (_| | | | (_) | |_| | | | | (_| |")
	fmt.Println(" \\_____\\__,_|___/\\__\\___|_|_|\\__,_|_| |_| \\___|  \\____/|_| |_|\\__,_|\\___|_|  \\__, |_|  \\___/ \\__,_|_| |_|\\__,_|")
	fmt.Println("                                                                              __/ |                            ")
	fmt.Println("                                                                             |___/                             ")
	str := "L'histoire commence dans la cour d'un HLM avec un daron, le tonton ainsi qu'une daronne, La discution porte autour des resaux au sein du quartier, les trois en ont marre et decident de se révolter ! "
	for _, char := range str {
		fmt.Printf("%c", char)
		time.Sleep(20 * time.Millisecond)
	}
}

=======
>>>>>>> f13e4f7e053c42328883fddd5af8e8285a81a77c
func ChoixPersonnage() character.Personnage {
	time.Sleep(3 * time.Second)
	fmt.Println("\n\nDARON : 55 ans, fumeur, passe son temps au PMU, supporte l'OM plus que sa femme. 100 kg au compteur, se bat avec une ceinture. \n(Beaucoup de points de vie, dégâts faibles) TAPEZ 1\n\nDARONNE : 44 ans, femme de ménage, regarde Les Marseillais et TPMP tous les soirs, brune, 60 kg, lance des claquettes pour attaquer. \n(Dégâts élevés, peu de points de vie) TAPEZ 2\n\nTONTON : Tonton cool, il aime les femmes, la plage, les paris sportifs et la pétanque par-dessus tout, gagne sa vie au black et avec l'achat-revente de T-MAX volés,\nrêve d'une carrière de bouliste et ne se bat QU'AVEC SES BOULES UNIQUEMENT. Il est mauvais, mais avec un peu de chance, il peut infliger des dégâts énormes. \n(Points de vie moyens, dégâts aléatoires) TAPEZ 3")
	var choix_perso int
	fmt.Scan(&choix_perso)

	for choix_perso != 1 && choix_perso != 2 && choix_perso != 3 {
		fmt.Println("Entrez seulement 1, 2 ou 3. \n")
		fmt.Scan(&choix_perso)
	}
	var nameuser string
	var perso character.Personnage
	switch choix_perso {
	case 1:
		fmt.Println("\nChoisisez un pseudo pour votre personnage \n")
		fmt.Scan(&nameuser)
		for len(nameuser) > 12 {
			fmt.Println("\nEntrez maximum 12 caractère !\n")
			fmt.Scan(&nameuser)
		}
		perso = character.Personnage{"Daron", nameuser, 120, 10, 20, 120}
		fmt.Printf("Voici les stats du %v %v \nPoints de vie : %v\nArgent : %v", perso.Name, perso.NameUser, perso.Hp, perso.Gold)
	case 2:
		fmt.Println("\nChoisisez un pseudo pour votre personnage \n")
		fmt.Scan(&nameuser)
		for len(nameuser) > 12 {
			fmt.Println("\nEntrez maximum 12 caractère !\n")
			fmt.Scan(&nameuser)
		}
		perso = character.Personnage{"Daronne", nameuser, 80, 10, 35, 120}
		fmt.Printf("Voici les stats de la %v %v \nPoints de vie : %v\nArgent : %v", perso.Name, perso.NameUser, perso.Hp, perso.Gold)
	case 3:
		fmt.Println("\nChoisisez un pseudo pour votre personnage \n")
		fmt.Scan(&nameuser)
		for len(nameuser) > 12 {
			fmt.Println("\nEntrez maximum 12 caractère !\n")
			fmt.Scan(&nameuser)
		}
		rand.Seed(time.Now().UnixNano())
		argentrandom := rand.Intn(46) + 5
		perso = character.Personnage{"Tonton", nameuser, 100, argentrandom, 20, 120}
		fmt.Printf("\nVoici les stats du %v %v \nPoints de vie : %v\nArgent : %v\n", perso.Name, perso.NameUser, perso.Hp, perso.Gold)
	default:
	}
	return perso

}

func Firstfight(perso character.Personnage) {
	attack := 0
	guetteur := character.Enemy{"Guetteur", 100, 10}
	fmt.Println(`
   O                         O
  /|\                       /|\
  / \                       / \
`)
	for guetteur.Hp > 0 && perso.Hp > 0 {

		fmt.Println(perso.NameUser, "a", perso.Hp, "point de vie et le", guetteur.Name, "a", guetteur.Hp, "point de vie")
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
				fmt.Println("Vous attaquez le guetteur et vous lui infligez", perso.Damage, "points de dégât\n")
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
