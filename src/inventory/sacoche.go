package inventory

import (
	"fmt"
	"projet-red/object"
)

type Inventory struct {
	Weapon []object.Weapon
	Drink  []object.Drink
}

func test() {
	a := []int{1, 2, 3, 4, 5, 6}
	a = append(a[:2], a[3:]...)
	fmt.Printf(a)
}
