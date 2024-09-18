package fight

import (
	"fmt"
	"projet-red/character"
	"projet-red/game"
	"projet-red/inventory"
	"projet-red/menu"
	"projet-red/object"
	"time"
)

func ThirdFight(perso *character.Personnage, inv inventory.Inventory) inventory.Inventory {
	attack := 0
	Gofasteur := character.Enemy{Name: "Go fasteur", Hp: 300, Damage: 20}
	tissu := object.ObjectStats{Name: "tissu", Type: "Utilitaire", Damage: 0}
	fmt.Println("\nVous entrez dans un combat avec un Go fasteur !")
	fmt.Println(`
	   O                         O
	  /|\                       /|\
	  / \                       / \`)
	time.Sleep(2 * time.Second)

	for Gofasteur.Hp > 0 && perso.Hp > 0 {
		fmt.Println("\n--- Combat ---")

		game.DisplayHealth(perso.NameUser, perso.Hp, perso.Hpmax)
		game.DisplayHealth(Gofasteur.Name, Gofasteur.Hp, 100)
		menu.AfficherInventaireEnCombat(&inv)

		x := len(inv.SacocheCp)
		armeUtilisable := 0
		for i := 0; i < x; i++ {
			if inv.SacocheCp[i].Type != "Utilitaire" {
				armeUtilisable++
				if inv.SacocheCp[i].Type == "Soin" {
					fmt.Printf("\nPour utilisez le %v et soigner %v pv. Appuyez sur %v\n", inv.SacocheCp[i].Name, inv.SacocheCp[i].Damage, i)
				}
				if inv.SacocheCp[i].Type == "Damage" {
					fmt.Printf("\nPour utilisez le %v et infliger %v deg. Appuyez sur %v\n", inv.SacocheCp[i].Name, inv.SacocheCp[i].Damage, i)
				}
			}
		}

		fmt.Scan(&attack)

		for attack >= 0 && attack <= armeUtilisable {
			fmt.Printf("Entrez une option valide (1)\n\n")
			fmt.Scan(&attack)
		}

		switch attack {
		case 0:
			damage := perso.Damage
			Gofasteur.Hp -= damage
			fmt.Printf("Vous infligez %d points de dégât.\n", damage)
			time.Sleep(2 * time.Second)
			fmt.Println("Le guetteur riposte !")
			time.Sleep(1 * time.Second)
			perso.Hp -= Gofasteur.Damage
			fmt.Printf("Le Go fasteur vous inflige %d points de dégât.\n", Gofasteur.Damage)
			time.Sleep(2 * time.Second)

			if perso.Hp <= 0 {
				fmt.Println("\nVous êtes tombé au combat...")
			} else if Gofasteur.Hp <= 0 {
				fmt.Println("\nVous avez vaincu le Go fasteur ! Vous trouvez du tissu et 500 euros !!")
				inv.AddObject(tissu)
				perso.Gold += 500
			}
		}

	}
	return inv
}
