package jeutel

import (
	"fmt"
)

var board [3][3]string // Le plateau de jeu 3x3
var currentPlayer string

// Fonction pour initialiser le plateau
func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}
	currentPlayer = "X"
}

// Fonction pour afficher le plateau
func displayBoard() {
	fmt.Println("  0   1   2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d %s | %s | %s\n", i, board[i][0], board[i][1], board[i][2])
		if i < 2 {
			fmt.Println(" ---|---|---")
		}
	}
}

// Fonction pour jouer un coup
func makeMove() {
	var row, col int
	for {
		fmt.Printf("Joueur %s, entrez la ligne et la colonne (ex : 1 2) : ", currentPlayer)
		fmt.Scan(&row, &col)

		if row >= 0 && row < 3 && col >= 0 && col < 3 && board[row][col] == " " {
			board[row][col] = currentPlayer
			break
		} else {
			fmt.Println("Mouvement invalide. Essayez encore.")
		}
	}
}

// Vérifie si un joueur a gagné
func checkWinner() bool {
	// Vérifie les lignes, colonnes et diagonales
	for i := 0; i < 3; i++ {
		if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
			return true
		}
		if board[0][i] == currentPlayer && board[1][i] == currentPlayer && board[2][i] == currentPlayer {
			return true
		}
	}
	if board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer {
		return true
	}
	if board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer {
		return true
	}
	return false
}

// Vérifie si le plateau est plein
func isBoardFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}

// Fonction principale pour jouer au morpion
func PlayMorpion() {
	initializeBoard()

	for {
		displayBoard()
		makeMove()

		if checkWinner() {
			displayBoard()
			fmt.Printf("Félicitations ! Le joueur %s a gagné !\n", currentPlayer)
			break
		}

		if isBoardFull() {
			displayBoard()
			fmt.Println("Match nul ! Le plateau est plein.")
			break
		}

		// Alterne entre les joueurs
		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}
}
