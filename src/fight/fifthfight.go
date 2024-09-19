package fight

import (
	"fmt"
	"projet-red/character"
	"projet-red/game"
	"projet-red/inventory"
	"time"
)

var Gérant = character.Enemy{Name: "Gérant", Hp: 300, Damage: 25}
var Guetteur = character.Enemy{Name: "Guetteur 1", Hp: 100, Damage: 10}
var Guetteur2 = character.Enemy{Name: "Guetteur 2", Hp: 100, Damage: 10}

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
				if inv.SacocheCp[i].Name == "Redbull" {
					fmt.Printf("\nPour utilisez le %v et augmenter vos dégâts %v%%. Appuyez sur %v\n", inv.SacocheCp[i].Name, inv.SacocheCp[i].Damage, i)
				}
				fmt.Printf("\nPour utiliser le %v et soigner %v pv. Appuyez sur %v\n", inv.SacocheCp[i].Name, inv.SacocheCp[i].Damage, i)
			}
			if inv.SacocheCp[i].Type == "Arme" {
				if inv.SacocheCp[i].Name == "Boule" {
					fmt.Printf("\nPour utiliser la %v et infliger des dégâts aléatoire. Appuyez sur %v\n", inv.SacocheCp[i].Name, i)
				} else {
					fmt.Printf("\nPour utiliser le %v et infliger %v dégâts. Appuyez sur %v\n", inv.SacocheCp[i].Name, inv.SacocheCp[i].Damage, i)
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

	if item.Name == "Redbull" {
		// Gestion du Redbull qui augmente les dégâts pendant 3 tours
		fmt.Println("Vous buvez un Redbull, vos dégâts sont augmentés pendant 3 tours !")
		perso.BoostDamageTurns = 4
		inv.RemoveObject(item) // Supprimer l'objet après usage
		return
	}

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

		switch attack2 {
		case 1:
			if enemy.Hp > 0 {
				// Gestion du Taser
				if item.Name == "Taser" {
					enemy.Immobilized = true
					enemy.ImmobilizedTurns = 1
					fmt.Printf("Vous utilisez le Taser sur %s. Il est immobilisé pour 1 tour !\n", enemy.Name)
					inv.RemoveObject(item)
					return
				}
				// Gestion de la Lacrymogène
				if item.Name == "Lacrymogène" {
					enemy.LacrymogèneActive = true
					enemy.LacrymogèneTurns = 3
					fmt.Printf("Vous lancez une lacrymogène sur %s. Il subira 15 points de dégâts par tour pendant 3 tours !\n", enemy.Name)
					inv.RemoveObject(item) // Suppression de l'objet après utilisation
					return
				}
				// Autres armes
				damage := item.Damage
				damageFloat := float32(damage) * 1.2
				if perso.BoostDamageTurns > 0 {
					damage = int(damageFloat)
				}
				enemy.Hp -= damage
				if enemy.Hp < 0 {
					enemy.Hp = 0
				}
				fmt.Printf("Vous infligez %d points de dégât à %s.\n", damage, enemy.Name)
			} else {
				fmt.Println("Le Gérant est déjà vaincu !")
			}

		case 2:
			if Guetteur.Hp > 0 {
				if item.Name == "Taser" {
					Guetteur.Immobilized = true
					Guetteur.ImmobilizedTurns = 1
					fmt.Printf("Vous utilisez le Taser sur %s. Il est immobilisé pour 1 tour !\n", Guetteur.Name)
					inv.RemoveObject(item)
					return
				}
				if item.Name == "Lacrymogène" {
					Guetteur.LacrymogèneActive = true
					Guetteur.LacrymogèneTurns = 3
					fmt.Printf("Vous lancez une lacrymogène sur %s. Il subira 15 points de dégâts par tour pendant 3 tours !\n", Guetteur.Name)
					inv.RemoveObject(item)
					return
				}
				damage := item.Damage
				if perso.BoostDamageTurns > 0 {
					damage += 10
				}
				Guetteur.Hp -= damage
				if Guetteur.Hp < 0 {
					Guetteur.Hp = 0
				}
				fmt.Printf("Vous infligez %d points de dégât à %s.\n", damage, Guetteur.Name)
			} else {
				fmt.Println("Le Guetteur 1 est déjà vaincu !")
			}

		case 3:
			if Guetteur2.Hp > 0 {
				if item.Name == "Taser" {
					Guetteur2.Immobilized = true
					Guetteur2.ImmobilizedTurns = 2
					inv.RemoveObject(item)
					fmt.Printf("Vous utilisez le Taser sur %s. Il est immobilisé pour 1 tour !\n", Guetteur2.Name)
					inv.RemoveObject(item)
					return
				}
				if item.Name == "Lacrymogène" {
					Guetteur2.LacrymogèneActive = true
					Guetteur2.LacrymogèneTurns = 3
					fmt.Printf("Vous lancez une lacrymogène sur %s. Il subira 15 points de dégâts par tour pendant 3 tours !\n", Guetteur2.Name)
					inv.RemoveObject(item)
					return
				}
				damage := item.Damage
				if perso.BoostDamageTurns > 0 {
					damage += 10
				}
				Guetteur2.Hp -= damage
				if Guetteur2.Hp < 0 {
					Guetteur2.Hp = 0
				}
				fmt.Printf("Vous infligez %d points de dégât à %s.\n", damage, Guetteur2.Name)
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

	// Réduire le nombre de tours pour le boost de dégâts du joueur (Redbull)
	if perso.BoostDamageTurns > 0 {
		perso.BoostDamageTurns--
		if perso.BoostDamageTurns == 0 {
			fmt.Println("L'effet de Redbull est terminé, vos dégâts reviennent à la normale.")
		}
	}

	// Gérer l'immobilisation des ennemis
	if enemy.Immobilized {
		if enemy.ImmobilizedTurns == 1 {
			fmt.Printf("%s est immobilisé et ne peut pas attaquer ce tour.\n", enemy.Name)
		}
		if enemy.ImmobilizedTurns <= 0 {
			enemy.Immobilized = false
			fmt.Printf("%s n'est plus immobilisé.\n", enemy.Name)
		}
	}

	if guetteur1.Immobilized {
		if guetteur1.ImmobilizedTurns == 1 {
			fmt.Printf("%s est immobilisé et ne peut pas attaquer ce tour.\n", guetteur2.Name)
		}
		if guetteur1.ImmobilizedTurns <= 0 {
			guetteur1.Immobilized = false
			fmt.Printf("%s n'est plus immobilisé.\n", guetteur1.Name)
		}
	}

	if guetteur2.Immobilized {
		guetteur2.ImmobilizedTurns--
		if guetteur2.ImmobilizedTurns == 1 {
			fmt.Printf("%s est immobilisé et ne peut pas attaquer ce tour.\n", guetteur2.Name)
		}
		if guetteur2.ImmobilizedTurns <= 0 {
			guetteur2.Immobilized = false
			fmt.Printf("%s n'est plus immobilisé.\n", guetteur2.Name)
		}
	}

	// Gérer la lacrymogène
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

	// Riposte des ennemis si non immobilisés
	totalDamage := 0

	if guetteur1.Hp <= 0 && guetteur2.Hp > 0 {
		if !guetteur2.Immobilized && !enemy.Immobilized {
			totalDamage = enemy.Damage + guetteur2.Damage
			fmt.Printf("Le Gérant et le guetteur 2 vous infligent %d points de dégât.\n", totalDamage)
		} else if enemy.Immobilized {
			totalDamage = guetteur2.Damage
			fmt.Printf("Le Gérant est immobilisé et le guetteur 2 vous inflige %d points de dégât.\n", totalDamage)
		} else if guetteur2.Immobilized {
			totalDamage = enemy.Damage
			fmt.Printf("Le guetteur 2 est immobilisé et le gérant vous inflige %d points de dégât.\n", totalDamage)
		}

	} else if guetteur2.Hp <= 0 && guetteur1.Hp > 0 {
		if !guetteur1.Immobilized && !enemy.Immobilized {
			totalDamage = enemy.Damage + guetteur1.Damage
			fmt.Printf("Le Gérant et le guetteur 1 vous infligent %d points de dégât.\n", totalDamage)
		} else if enemy.Immobilized {
			totalDamage = guetteur1.Damage
			fmt.Printf("Le Gérant est immobilisé et le guetteur 2 vous inflige %d points de dégât.\n", totalDamage)
		} else if guetteur1.Immobilized {
			totalDamage = enemy.Damage
			fmt.Printf("Le guetteur 1 est immobilisé et le Gérant vous inflige %d points de dégât.\n", totalDamage)
		}

	} else if guetteur1.Hp <= 0 && guetteur2.Hp <= 0 {
		if !enemy.Immobilized {
			totalDamage = enemy.Damage
			fmt.Printf("Le Gérant vous inflige %d points de dégât.\n", totalDamage)
		} else if enemy.Immobilized {
			fmt.Printf("Le Gérant est immobilisé et il ne vous inflige aucun dégât.\n")
		}

	} else if enemy.Hp <= 0 {
		if !guetteur1.Immobilized && !guetteur2.Immobilized {
			totalDamage = guetteur1.Damage + guetteur2.Damage
			fmt.Printf("Les guetteurs vous infligent %d points de dégât.\n", totalDamage)
		} else if guetteur1.Immobilized {
			totalDamage = guetteur2.Damage
			fmt.Printf("Le guetteur 1 est immobilisé et le guetteur 2 vous inflige %d points de dégât.\n", totalDamage)
		} else if guetteur2.Immobilized {
			totalDamage = guetteur1.Damage
			fmt.Printf("Le guetteur 2 est immobilisé et le guetteur 1 vous inflige %d points de dégât.\n", totalDamage)
		}
	} else {
		totalDamage = enemy.Damage + guetteur1.Damage + guetteur2.Damage
		fmt.Printf("Le Gérant et ses guetteurs vous infligent %d points de dégât.\n", totalDamage)
	}

	perso.Hp -= totalDamage

	time.Sleep(2 * time.Second)
}
