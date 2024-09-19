package story

import (
	"fmt"
	"projet-red/ascii"
	"time"
)

func HistoireDebut() {
	fmt.Println("   _____          _       _ _                     _    _           _                                          _ ")
	fmt.Println("  / ____|        | |     | | |                   | |  | |         | |                                        | |")
	fmt.Println(" | |     __ _ ___| |_ ___| | | __ _ _ __   ___   | |  | |_ __   __| | ___ _ __ __ _ _ __ ___  _   _ _ __   __| |")
	fmt.Println(" | |    / _` / __| __/ _ | | |/ _` | '_ \\ / _ \\  | |  | | '_ \\ / _` |/ _ | '__/ _` | '__/ _ \\| | | | '_ \\ / _` |")
	fmt.Println(" | |___| (_| \\__ | ||  __| | | (_| | | | |  __/  | |__| | | | | (_| |  __| | | (_| | | | (_) | |_| | | | | (_| |")
	fmt.Println(" \\_____\\__,_|___/\\__\\___|_|_|\\__,_|_| |_| \\___|  \\____/|_| |_|\\__,_|\\___|_|  \\__, |_|  \\___/ \\__,_|_| |_|\\__,_|")
	fmt.Println("                                                                              __/ |                            ")
	ascii.Jul()
	fmt.Println("                                                                             |___/                             ")
	str := "L’histoire commence dans la cour d’un HLM avec un daron, un tonton ainsi qu’une daronne. La discussion porte autour des réseaux au sein du quartier,\nles trois en ont marre et décident de se révolter ! Choisis ton personnage !"
	for _, char := range str {
		fmt.Printf("%c", char)
		time.Sleep(40 * time.Millisecond)
	}
}

func Afterchoixperso() {
	str := "\nTrès bien ! La discussion prend fin et tu décides de te rendre à l’entrée du quartier afin de rencontrer le guetteur pour récolter des informations sur ses supérieurs,\nmais celui-ci n’est pas ouvert au débat. Le combat commence !"
	for _, char := range str {
		fmt.Printf("%c", char)
		time.Sleep(40 * time.Millisecond)
	}
}

func Afterguetteur() {
	str := "\nAdversaire vaincu ! Vous vous en sortez de justesse ! Tu récupères sa sacoche C.P. Company contenant : Un briquet, 10 €, Un sandwich, 2 Ricards ainsi qu’un téléphone\n(accès à Telegram, le shop du jeu, besoin d’un item ? vas y faire un tour. "
	for _, char := range str {
		fmt.Printf("%c", char)
		time.Sleep(40 * time.Millisecond)
	}
}

func AfterVendeur() {
	str := "\n Apres quelques minutes de marche tu croise un RS6 gris nardo plein a craquer et tu decides de te battre avec le chauffeur de ce go-Fast (Tu attaque en premier, le chauffeur se bat avec un couteau suisse trouvé dans la boite à gants "
	for _, char := range str {
		fmt.Printf("%c", char)
		time.Sleep(40 * time.Millisecond)
	}
}
