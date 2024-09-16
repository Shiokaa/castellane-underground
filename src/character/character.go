package character

type Enemy struct {
	Name   string
	Hp     int
	Damage int
}

type Personnage struct {
	Name     string
	NameUser string
	Hp       int
	Gold     int
	Damage   int
	Hpmax    int
}
