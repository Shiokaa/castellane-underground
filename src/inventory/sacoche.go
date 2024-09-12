package inventory

import (
	"fmt"
	"projet-red/object"
)

type Inventory struct {
	SacocheCp []object.Object
	Limite    int
}

func test() {
	a := []int{1, 2, 3, 4, 5, 6}
	a = append(a[:2], a[3:]...)
	fmt.Printf(a)
}
