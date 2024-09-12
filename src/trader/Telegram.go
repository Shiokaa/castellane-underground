package trader

import (
	"fmt"
	"projet-red/object"
)

func Telegram() {
	inv := []object.Object{}
	achat := 0
	fmt.Println("Tapez 1 pour acheter une lacrimogène 	Tapez 5 pour avoir un ricard \nTapez 2 pour acheter la matraque 	Tapez 6 pour avoir un flash \nTapez 3 pour acheter un mortier 	Tapez 7 pour avoir une redbull \nTapez 4 pour acheter un taser      	Tapez 8 pour avoir la sacoche lv")
	fmt.Scan(&achat)
	for achat != 1 && achat != 2 && achat != 3 && achat != 4 && achat != 5 && achat != 6 && achat != 7 && achat != 8 {
		fmt.Println("Entrez un numéro valide\n")
		fmt.Scan(&achat)
	}
	switch achat {
	case 1:
		inv = append(inv, object.Object{Lacrimogene: object.ObjectStats{Name: "Lacrimogène"}})
		fmt.Println("Une lacrimogène a été ajoutée à l'inventaire")
	case 2:
		inv = append(inv, object.Object{Matraque: object.ObjectStats{Name: "Matraque"}})
		fmt.Print("La matraque a été ajoutée à l'inventaire")
	case 3:
		inv = append(inv, object.Object{Mortier: object.ObjectStats{Name: "Mortier"}})
		fmt.Println("Un mortier a été ajoutée à l'inventaire")
	case 4:
		inv = append(inv, object.Object{Taser: object.ObjectStats{Name: "Taser"}})
		fmt.Print("Un taser a été ajoutée à l'inventaire")
	case 5:
		inv = append(inv, object.Object{Ricard: object.ObjectStats{Name: "Ricard"}})
		fmt.Println("Un ricard a été ajoutée à l'inventaire")
	case 6:
		inv = append(inv, object.Object{Flash: object.ObjectStats{Name: "Flash"}})
		fmt.Print("Un flash a été ajoutée à l'inventaire")
	case 7:
		inv = append(inv, object.Object{Redbull: object.ObjectStats{Name: "Redbull"}})
		fmt.Print("Une redbull a été ajoutée à l'inventaire")
	case 8:
		fmt.Println("Vous avez acheté la sacoche lv vous gagnez 10 emplacement")
	default:
	}
}
