package character

type Personnage struct {
	Name     string
	NameUser string
	Hp       int
	Gold     int
	Damage   int
	Hpmax    int
}

type Enemy struct {
	Name   string
	Hp     int
	Damage int
}
