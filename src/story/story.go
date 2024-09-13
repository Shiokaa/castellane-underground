package story

import (
	"fmt"
	"time"
)

func HistoireDebut() {
	str := "L'histoire commence dans la cour d'un HLM avec un daron, le tonton ainsi qu'une daronne. La discussion porte autour des réseaux au sein du quartier, les trois en ont marre et \ndécident de se révolter !"
	for _, char := range str {
		fmt.Printf("%c", char)
		time.Sleep(40 * time.Millisecond)
	}
}

func Afterchoixperso() {
	str := "\nVous vous baladez dans le quartier quand vous croisez un guetteur. Vous décidez de commencer votre rébellion."
	for _, char := range str {
		fmt.Printf("%c", char)
		time.Sleep(40 * time.Millisecond)
	}
}

func Afterguetteur() {
	str := "\nVous avez tué le guetteur, vous recevez sa sacoche CP, son téléphone et vous débloquez l'accès à l'inventaire."
	for _, char := range str {
		fmt.Printf("%c", char)
		time.Sleep(40 * time.Millisecond)
	}
}
