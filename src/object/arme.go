package object

// Structure générique pour chaque arme
type WeaponStats struct {
    Name   string // Nom de l'arme
    Damage int    // Dégâts infligés
}

// Structure principale pour les armes
type Weapon struct {
    Ceinture    WeaponStats // Ceinture (-20)
    Claquette   WeaponStats // Claquette (-30)
    Boule       WeaponStats // Boule (-5 ~ -200)
    Lacrimogene WeaponStats // Lacrimogène (-5% hp pendant 3 tours)
    Matraque    WeaponStats // Matraque (-80)
    Cocktail    WeaponStats // Cocktail Molotov (-75 pdt 3 tours)
    Taser       WeaponStats // Taser (-100 + stun pendant 1 tour)
    Mortier     WeaponStats // Mortier (-200)
	Poing       WeaponStats // Poing (-10)
}
