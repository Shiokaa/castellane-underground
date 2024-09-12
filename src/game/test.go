package game

import (
	"fmt"
	"projet-red/character"
	"projet-red/object"
)

func Test() {
	attaquer := 0
	guetteur := character.Enemy{Hp: 100}
	poigdmg := object.Weapon{Poing: 10}
	fmt.Println("Attaquer pressez 1")
	fmt.Scan(&attaquer)
	if attaquer == 1 {
		guetteur.Hp -= poigdmg.Poing
	}
	fmt.Println(guetteur.Hp)
}
