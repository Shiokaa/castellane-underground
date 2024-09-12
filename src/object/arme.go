package object

// Structure générique pour chaque arme
type ObjectStats struct {
	Name   string // Nom de l'arme
	Type   string //soin ou arme
	Damage int    // Dégâts infligés ou soins recu
}

// Structure principale pour les armes
type Object struct {
	Ceinture    ObjectStats // Ceinture (-20)
	Claquette   ObjectStats // Claquette (-30)
	Boule       ObjectStats // Boule (-5 ~ -200)
	Lacrimogene ObjectStats // Lacrimogène (-5% hp pendant 3 tours)
	Matraque    ObjectStats // Matraque (-80)
	Cocktail    ObjectStats // Cocktail Molotov (-75 pdt 3 tours)
	Taser       ObjectStats // Taser (-100 + stun pendant 1 tour)
	Mortier     ObjectStats // Mortier (-200)
	Poing       ObjectStats // Poing (-10)

	Ricard    ObjectStats // Ricard
	Flash     ObjectStats // Flash
	Sandwitch ObjectStats // Sandwitch
	Redbull   ObjectStats // Redbull
}
