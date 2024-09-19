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
	fmt.Println("                                                                             |___/                             ")
	fmt.Println("\033[34m")
	ascii.OM()
	fmt.Println("\033[0m")
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
	str := "\nTu te balades tranquillement, tout à coup, tu aperçois une Audi RS6 gris Nardo, pleine à craquer, passer à fond à côté de toi et se garer 10 mètres plus loin. Tu décides d’aller voir ce qui se cache à l’intérieur mais le chauffeur du go-fast n’est absolument pas d’accord a l’idée de regarder ce qu’il s’y cache."
	for _, char := range str {
		fmt.Printf("%c", char)
		time.Sleep(40 * time.Millisecond)
	}
}

func AfterGoFast() {
	str := "\n. Au loin, l’homme de main t’a repéré (la brute du quartier) et ne souhaite pas t’épargner, le combat commence !"
	for _, char := range str {
		fmt.Printf("%c", char)
		time.Sleep(40 * time.Millisecond)
	}
}
