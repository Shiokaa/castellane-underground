package object

type DrinkStats struct {
	Name  string // Nom de la boisson
	Value int    // Valeur liée à l'effet de la boisson
}

// Structure principale pour les boissons
type Drink struct {
	Ricard    DrinkStats // Ricard
	Flash     DrinkStats // Flash
	Sandwitch DrinkStats // Sandwitch
	Redbull   DrinkStats // Redbull
}
