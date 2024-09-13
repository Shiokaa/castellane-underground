package inventory

import (
	"projet-red/object"
)

type Inventory struct {
	SacocheCp      []object.ObjectStats
	CraftInventory []object.ObjectStats
	Limite         int
}
