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

func (inv *Inventory) RemoveObject(obj object.ObjectStats) {
	for index, elem := range inv.SacocheCp {
		if elem == obj {
			inv.SacocheCp = append(inv.SacocheCp[:index], inv.SacocheCp[index+1:]...)
		}
	}
}
