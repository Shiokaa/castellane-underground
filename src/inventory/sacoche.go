package inventory

import (
	"fmt"
	"projet-red/object"
)

type Inventory struct {
	Weapon []object.Weapon
	Drink  []object.Drink
}

func (inv *Inventory) AjouterWeapon(object.Weapon) {
	inv.Weapon = append(inv.Weapon, object.Weapon)
}
func (inv *Inventory) AjouterDrink(object.drink) {
	inv.Drink = append(inv.Drink, object.Drink)
}

func test() {
	a := []int{1, 2, 3, 4, 5, 6}
	a = append(a[:2], a[3:]...)
	fmt.Printf(a)
}
