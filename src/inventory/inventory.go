package inventory

import (
	"projet-red/object"
)

type Inventory struct {
	SacocheCp      [8]object.ObjectStats
	CraftInventory []object.ObjectStats
	Limite         int
}

func (inv *Inventory) AddObject(obj object.ObjectStats) {
	inv.SacocheCp = append(inv.SacocheCp, obj)
}
