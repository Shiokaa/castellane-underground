package character

import (
	"math/rand"
	"time"
)

type Personnage struct {
	Name           string
	NameUser       string
	Hp             int
	Gold           int
	Damage         int
	Hpmax          int
	CombatCounteur int
}

type Enemy struct {
	Name   string
	Hp     int
	Damage int
}

func DegatTonton() int {
	// Initialisation du générateur de nombres aléatoires avec un seed basé sur le temps
	rand.Seed(time.Now().UnixNano())

	// Génération d'un nombre aléatoire entre 5 et 200
	min := 5
	max := 200
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber
}
