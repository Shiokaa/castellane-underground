# castellane-underground

### Description
Castellane Underground est un jeu de rôle textuel développé en Go, inspiré par les jeux Pokémon classiques, mais se déroulant dans le quartier de Castellane à Marseille. Le joueur incarne un personnage qui explore le quartier, combat divers adversaires, et utilise des ressources pour avancer dans l'histoire. Le jeu propose un système de combat tour par tour, une économie, et des options de fabrication d'armes, le tout dans une ambiance humoristique inspirée de la culture marseillaise.

### Caractéristiques

-Création de personnage : Personnalisez votre personnage et plongez dans l'univers de la Castellane.
-Système de combat : Combattez vos adversaires en utilisant un système de combat au tour par tour.
-Économie et inventaire : Gérer vos ressources, acheter et utiliser des objets, ainsi que des armes à dégâts temporaires.
-Fabrication d'armes : Rassemblez des ressources pour fabriquer une arme mystère.
-Scénario interactif : Prenez des décisions qui influenceront le déroulement de l'histoire.

### Installation

#### prérequis
Go version 1.17 ou plus récente doit être installé sur votre machine.

Cloner le dépots :
git clone https://github.com/Shiokaa/castellane-underground.git


### Utilisation
Après avoir lancé le jeu, vous serez invité à créer votre personnage et à choisir un nom. Suivez les instructions à l'écran pour naviguer dans le quartier, interagir avec les personnages et participer à des combats. Le but est d'explorer Castellane, de survivre aux combats, et de progresser dans l'histoire.

``` castellane_underground/
├── main.go            # Point d'entrée du programme
├── game/
│   ├── combat.go      # Gestion du système de combat
│   ├── economy.go     # Gestion de l'économie et de l'inventaire
│   ├── player.go      # Définition du personnage et des actions du joueur
│   └── story.go       # Scénario et dialogues du jeu
├── assets/            # Contient les ressources du jeu (fichiers texte, ASCII art)
└── README.md          # Documentation du projet ```


