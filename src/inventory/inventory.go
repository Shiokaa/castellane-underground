package inventory

import (
	"projet-red/object"
)

type Inventory struct {
	SacocheCp      []object.ObjectStats
	CraftInventory []object.ObjectStats
	Limite         int
}

func (inv *Inventory) AddObject(obj object.ObjectStats) {
	inv.SacocheCp = append(inv.SacocheCp, obj)
}

func (inv *Inventory) RemoveObject(object.ObjectStats) {
	for index, elem := range inv.SacocheCp {
		if elem.Name == "" {
			if index == len(inv.SacocheCp)-1 {
				inv.SacocheCp = inv.SacocheCp[:index]
			} else {
				inv.SacocheCp = append(inv.SacocheCp[:index], inv.SacocheCp[index+1:]...)
			}
		}
	}
}
