package inventory

import (
	"fmt"
	"projet-red/object"

	"time"
)

type Inventory struct {
	SacocheCp      []object.ObjectStats
	CraftInventory []object.ObjectStats
	Limite         int
}

func (inv *Inventory) AddObject(obj object.ObjectStats) {
	inv.SacocheCp = append(inv.SacocheCp, obj)
}
func (inv *Inventory) AddCraft(obj object.ObjectStats) {
	inv.CraftInventory = append(inv.CraftInventory, obj)
}

func (inv *Inventory) AfficherInventaireEnCombat() {
	fmt.Println("\n--------- INVENTAIRE ---------")
	if len(inv.SacocheCp) > 0 {
		for i, obj := range inv.SacocheCp {
			fmt.Printf("%d. %s (%s)\n", i+1, obj.Name, obj.Type)
		}
	} else {
		fmt.Println("Votre inventaire est vide.")
	}
	fmt.Println("------------------------------")
	time.Sleep(2 * time.Second)
}

func (inv *Inventory) RemoveObject(object object.ObjectStats) {
	Nomobjet := object.Name
	for index := range inv.SacocheCp {
		if inv.SacocheCp[index].Name == Nomobjet {
			if index == len(inv.SacocheCp)-1 {
				inv.SacocheCp = inv.SacocheCp[:index]
			} else {
				inv.SacocheCp = append(inv.SacocheCp[:index], inv.SacocheCp[index+1:]...)
			}
			break
		}
	}
}
