package game

import (
	"fmt"
	"math/rand"
	"projet-red/character"
	"projet-red/inventory"
	"projet-red/object"
	"time"
)

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
		perso = character.Personnage{"Daronne", nameuser, 80, 10, 35, 80}
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
		perso = character.Personnage{"Tonton", nameuser, 100, argentrandom, 20, 100}
		fmt.Printf("\nVoici les stats du %v %v \nPoints de vie : %v\nArgent : %v\n", perso.Name, perso.NameUser, perso.Hp, perso.Gold)
	default:
	}
	return perso
}

// Fonction pour afficher la santé du personnage avec une barre dynamique
func displayHealth(entityName string, hp int, hpMax int) {
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

func Firstfight(perso *character.Personnage) inventory.Inventory {
	inv := inventory.Inventory{Limite: 5}
	briquet := object.ObjectStats{"Briquet", "Utilitaire", 10}
	sandwitch := object.ObjectStats{"Sandwitch", "Utilitaire", 10}
	ricard := object.ObjectStats{"Ricard", "Utilitaire", 10}

	attack := 0
	guetteur := character.Enemy{"Guetteur", 100, 10}

	// Combat Intro
	fmt.Println("\nVous entrez dans un combat avec un Guetteur !")
	fmt.Println(`
   O                         O
  /|\                       /|\
  / \                       / \
`)
	time.Sleep(2 * time.Second)

	// Boucle de combat
	for guetteur.Hp > 0 && perso.Hp > 0 {
		fmt.Println("\n--- Combat ---")

		// Utilisation de la nouvelle fonction displayHealth avec hpMax dynamique
		displayHealth(perso.NameUser, perso.Hp, perso.Hpmax)
		displayHealth(guetteur.Name, guetteur.Hp, 100)

		// Options de combat
		fmt.Println("\nQue voulez-vous faire ?")
		fmt.Println("1 - Attaquer")
		fmt.Scan(&attack)

		// Vérification de l'entrée
		for attack != 1 && attack != 2 {
			fmt.Println("Entrez une option valide (1)\n")
			fmt.Scan(&attack)
		}

		// Gestion des actions
		switch attack {
		case 1: // Attaquer
			damage := perso.Damage
			guetteur.Hp -= damage
			fmt.Printf("Vous infligez %d points de dégât.\n", damage)
			time.Sleep(2 * time.Second)

			// Attaque du Guetteur
				if guetteur.Hp > 21 {
					fmt.Println("Le guetteur riposte !")
					time.Sleep(1 * time.Second)
					perso.Hp -= guetteur.Damage
					fmt.Printf("Le guetteur vous inflige %d points de dégât.\n", guetteur.Damage)
				} else {
					// Le guetteur sort un couteau si sa santé est basse
					criticalDamage := 50
					perso.Hp -= criticalDamage
					fmt.Println("Le guetteur sort un couteau !")
					fmt.Printf("Il vous inflige un coup critique de %d points de dégât.\n", criticalDamage)
				}
				time.Sleep(2 * time.Second)

			// Si le joueur ou le guetteur sont morts
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
