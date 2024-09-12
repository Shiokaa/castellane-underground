package trader

import (
	"fmt"
	"projet-red/object"
)

func Telegram() {
	inv := []object.Weapon{}
	achat := 0
	fmt.Println("Tapez 1 pour acheter la lacrimogène 	Tapez 5 pour avoir un ricard \nTapez 2 pour acheter la matraque 	Tapez 6 pour avoir un flash \nTapez 3 pour acheter un mortier 	Tapez 7 pour avoir une redbull \nTapez 4 pour acheter un taser")
	fmt.Scan(&achat)
	switch achat {
	case 1:
		inv = append(inv, object.Weapon{Lacrimogene: object.WeaponStats{Name: "Lacrimogène"}})
		fmt.Println("La matraque a été ajoutée à l'inventaire")
	case 2:
		inv = append(inv, object.Weapon{Matraque: object.WeaponStats{Name: "Matraque"}})
		fmt.Print("Le cocktail a été ajoutée à l'inventaire")
	case 3:
		inv = append(inv, object.Weapon{Mortier: object.WeaponStats{Name: "Mortier"}})
		fmt.Println("La matraque a été ajoutée à l'inventaire")
	case 4:
		inv = append(inv, object.Weapon{Taser: object.WeaponStats{Name: "Taser"}})
		fmt.Print("Le cocktail a été ajoutée à l'inventaire")
	case 5:
		inv = append(inv, object.Drink{Ricard: object.WeaponStats{Name: "Ricard"}})
		fmt.Println("La matraque a été ajoutée à l'inventaire")
	case 6:
		inv = append(inv, object.Weapon{Flash: object.WeaponStats{Name: "Flash"}})
		fmt.Print("Le cocktail a été ajoutée à l'inventaire")
	case 7:
		inv = append(inv, object.Weapon{Redbull: object.WeaponStats{Name: "Redbull"}})
		fmt.Print("Le cocktail a été ajoutée à l'inventaire")
	default:
	}
}
