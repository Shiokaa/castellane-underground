package character

type Enemy struct {
	name   string
	hp     int
	damage int
}

func init(){
	guetteur := Enemy{"Guetteur", 100 , 10}
	vendeur := Enemy{}
}