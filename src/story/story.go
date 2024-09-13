package story

import (
	"fmt"
	"time"
)

func HistoireDebut() {
	str := "L’histoire commence dans la cour d’un HLM avec un daron, un tonton ainsi qu’une daronne. La discussion porte autour des réseaux au sein du quartier,\n les trois en ont marre et décident de se révolter ! Choisis ton personnage !"
	for _, char := range str {
		fmt.Printf("%c", char)
		time.Sleep(40 * time.Millisecond)
	}
}

func Afterchoixperso() {
	str := "\nTrès bien ! La discussion prend fin et tu décides de te rendre à l’entrée du quartier afin de rencontrer le guetteur pour récolter des informations sur ses supérieurs,\n mais celui-ci n’est pas ouvert au débat. Le combat commence !"
	for _, char := range str {
		fmt.Printf("%c", char)
		time.Sleep(40 * time.Millisecond)
	}
}

func Afterguetteur() {
	str := "\nAdversaire vaincu ! Tu récupères sa sacoche C.P. Company contenant : Un briquet, 10 €, Un sandwich, 2 Ricards ainsi qu’un téléphone (accès à Telegram,\n le shop du jeu, besoin d’un item ? vas y faire un tour. "
	for _, char := range str {
		fmt.Printf("%c", char)
		time.Sleep(40 * time.Millisecond)
	}
}
