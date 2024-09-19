package fight

import (
	"fmt"
	"projet-red/character"
	"projet-red/game"
	"projet-red/inventory"
	"time"
)

var Gérant = character.Enemy{Name: "Gérant", Hp: 300, Damage: 25}
var Guetteur = character.Enemy{Name: "Guetteur", Hp: 100, Damage: 10}
var Guetteur2 = character.Enemy{Name: "Guetteur", Hp: 100, Damage: 10}

func FifthFight(perso *character.Personnage, inv inventory.Inventory) inventory.Inventory {
	fmt.Println("\nVous entrez dans un combat avec un Gérant !")
	fmt.Println(`
	   O                         O
	  /|\                       /|\
	  / \                       / \`)
	time.Sleep(2 * time.Second)

	for Gérant.Hp > 0 && perso.Hp > 0 {
		if perso.Name == "Tonton" {
			perso.Damage = character.DegatTonton()
		}
		fmt.Println("\n--- Combat ---")
		game.DisplayHealth(perso.NameUser, perso.Hp, perso.Hpmax)
		fmt.Println("")
		game.DisplayHealth(Gérant.Name, Gérant.Hp, 300)
		game.DisplayHealth(Guetteur.Name, Guetteur.Hp, 100)
		game.DisplayHealth(Guetteur2.Name, Guetteur2.Hp, 100)

		// Sélection de l'action par l'utilisateur
		attack := chooseAction2(len(inv.SacocheCp), inv)

		// Appliquer l'action choisie
		handleAction2(attack, &Gérant, perso, &inv)

		// Vérifier si l'un des deux personnages est mort
		if Gérant.Hp <= 0 {
			fmt.Println("\nVous avez vaincu le Gérant ! Vous trouvez 500 euros !!")
			if perso.CombatCounteur < 4 {
				perso.CombatCounteur = 4
			}
			perso.Gold += 500
			break
		} else if perso.Hp <= 0 {
			break
		}

		// Riposte du Gérant et des guetteurs
		enemyRetaliation2(&Gérant, perso, &Guetteur, &Guetteur2)
	}

	fmt.Println("\nVous êtes tombé au combat...")
	perso.Hp = perso.Hpmax / 2
	perso.Gold /= 2
	time.Sleep(5 * time.Second)
	return inv
}

func chooseAction2(max int, inv inventory.Inventory) int {
	var attack int
	for i := 0; i < len(inv.SacocheCp); i++ {
		if inv.SacocheCp[i].Type != "Utilitaire" {
			if inv.SacocheCp[i].Type == "Soin" {
				fmt.Printf("\nPour utilisez le %v et soigner %v pv. Appuyez sur %v\n", inv.SacocheCp[i].Name, inv.SacocheCp[i].Damage, i)
			}
			if inv.SacocheCp[i].Type == "Arme" {
				if inv.SacocheCp[i].Name == "Boule" {
					fmt.Printf("\nPour utilisez le %v et infliger des degats aléatoire. Appuyez sur %v\n", inv.SacocheCp[i].Name, i)
				} else {
					fmt.Printf("\nPour utilisez le %v et infliger %v dégât. Appuyez sur %v\n", inv.SacocheCp[i].Name, inv.SacocheCp[i].Damage, i)
				}
			}
		}
	}
	for {
		fmt.Print("\nEntrez votre choix : ")
		fmt.Scan(&attack)
		if attack >= 0 && attack < max {
			return attack
		}
		fmt.Println("Entrez une option valide.")
	}
}

func handleAction2(attack int, enemy *character.Enemy, perso *character.Personnage, inv *inventory.Inventory) {
	var attack2 int
	item := inv.SacocheCp[attack]

	if item.Type == "Arme" {
		fmt.Println("\nQui souhaitez-vous attaquer ?")

		// Afficher les options uniquement si l'ennemi est encore vivant
		if enemy.Hp > 0 {
			fmt.Println("1- Le Gérant")
		}
		if Guetteur.Hp > 0 {
			fmt.Println("2- Le Guetteur 1")
		}
		if Guetteur2.Hp > 0 {
			fmt.Println("3- Le Guetteur 2")
		}

		fmt.Scan(&attack2)

		// Vérification et application des dégâts
		switch attack2 {
		case 1:
			if enemy.Hp > 0 {
				if item.Name == "Lacrymogène" {
					enemy.LacrymogèneActive = true
					enemy.LacrymogèneTurns = 3
					inv.RemoveObject(item)
					fmt.Printf("Vous utilisez la lacrymogène sur %s. Il subira 15 points de dégât par tour pendant 3 tours.\n", enemy.Name)
				} else {
					damage := item.Damage
					enemy.Hp -= damage
					if enemy.Hp < 0 {
						enemy.Hp = 0
					}
					fmt.Printf("Vous infligez %d points de dégât à %s.\n", damage, enemy.Name)
				}
			} else {
				fmt.Println("Le Gérant est déjà vaincu !")
			}

		case 2:
			if Guetteur.Hp > 0 {
				if item.Name == "Lacrymogène" {
					Guetteur.LacrymogèneActive = true
					Guetteur.LacrymogèneTurns = 3
					inv.RemoveObject(item)
					fmt.Printf("Vous utilisez la lacrymogène sur %s. Il subira 15 points de dégât par tour pendant 3 tours.\n", Guetteur.Name)
				} else {
					damage := item.Damage
					Guetteur.Hp -= damage
					if Guetteur.Hp < 0 {
						Guetteur.Hp = 0
					}
					fmt.Printf("Vous infligez %d points de dégât à %s.\n", damage, Guetteur.Name)
				}
			} else {
				fmt.Println("Le Guetteur 1 est déjà vaincu !")
			}

		case 3:
			if Guetteur2.Hp > 0 {
				if item.Name == "Lacrymogène" {
					Guetteur2.LacrymogèneActive = true
					Guetteur2.LacrymogèneTurns = 3
					inv.RemoveObject(item)
					fmt.Printf("Vous utilisez la lacrymogène sur %s. Il subira 15 points de dégât par tour pendant 3 tours.\n", Guetteur2.Name)
				} else {
					damage := item.Damage
					Guetteur2.Hp -= damage
					if Guetteur2.Hp < 0 {
						Guetteur2.Hp = 0
					}
					fmt.Printf("Vous infligez %d points de dégât à %s.\n", damage, Guetteur2.Name)
				}
			} else {
				fmt.Println("Le Guetteur 2 est déjà vaincu !")
			}

		default:
			fmt.Println("Choix invalide. Veuillez sélectionner un ennemi vivant.")
		}

	} else if item.Type == "Soin" {
		heal := item.Damage
		perso.Hp += heal
		if perso.Hp > perso.Hpmax {
			perso.Hp = perso.Hpmax
		}
		fmt.Printf("Vous vous soignez de %d PV.\n", heal)

		inv.RemoveObject(item)
	}

	time.Sleep(2 * time.Second)
}

func enemyRetaliation2(enemy *character.Enemy, perso *character.Personnage, guetteur1 *character.Enemy, guetteur2 *character.Enemy) {
	time.Sleep(1 * time.Second)

	if enemy.LacrymogèneActive {
		enemy.Hp -= 15
		if enemy.Hp < 0 {
			enemy.Hp = 0
		}
		fmt.Printf("%s subit 15 points de dégâts à cause de la lacrymogène.\n", enemy.Name)
		enemy.LacrymogèneTurns--
		if enemy.LacrymogèneTurns <= 0 {
			enemy.LacrymogèneActive = false
			fmt.Printf("L'effet de la lacrymogène sur %s a pris fin.\n", enemy.Name)
		}
	}

	if guetteur1.LacrymogèneActive {
		guetteur1.Hp -= 15
		if guetteur1.Hp < 0 {
			guetteur1.Hp = 0
		}
		fmt.Printf("%s subit 15 points de dégâts à cause de la lacrymogène.\n", guetteur1.Name)
		guetteur1.LacrymogèneTurns--
		if guetteur1.LacrymogèneTurns <= 0 {
			guetteur1.LacrymogèneActive = false
			fmt.Printf("L'effet de la lacrymogène sur %s a pris fin.\n", guetteur1.Name)
		}
	}

	if guetteur2.LacrymogèneActive {
		guetteur2.Hp -= 15
		if guetteur2.Hp < 0 {
			guetteur2.Hp = 0
		}
		fmt.Printf("%s subit 15 points de dégâts à cause de la lacrymogène.\n", guetteur2.Name)
		guetteur2.LacrymogèneTurns--
		if guetteur2.LacrymogèneTurns <= 0 {
			guetteur2.LacrymogèneActive = false
			fmt.Printf("L'effet de la lacrymogène sur %s a pris fin.\n", guetteur2.Name)
		}
	}

	totalDamage := 0

	if guetteur1.Hp <= 0 && guetteur2.Hp > 0 {
		totalDamage = enemy.Damage + guetteur2.Damage
		fmt.Printf("Le Gérant et le guetteur 2 vous infligent %d points de dégât.\n", totalDamage)
	} else if guetteur2.Hp <= 0 && guetteur1.Hp > 0 {
		totalDamage = enemy.Damage + guetteur1.Damage
		fmt.Printf("Le Gérant et le guetteur 1 vous infligent %d points de dégât.\n", totalDamage)
	} else if guetteur1.Hp <= 0 && guetteur2.Hp <= 0 {
		totalDamage = enemy.Damage
		fmt.Printf("Le Gérant vous inflige %d points de dégât.\n", totalDamage)
	} else if enemy.Hp <= 0 {
		totalDamage = guetteur1.Damage + guetteur2.Damage
		fmt.Printf("Les guetteurs vous infligent %d points de dégât.\n", totalDamage)
	} else {
		totalDamage = enemy.Damage + guetteur1.Damage + guetteur2.Damage
		fmt.Printf("Le Gérant et ses guetteurs vous infligent %d points de dégât.\n", totalDamage)
	}

	perso.Hp -= totalDamage

	time.Sleep(2 * time.Second)
}
