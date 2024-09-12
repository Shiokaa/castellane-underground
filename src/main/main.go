package main

import "fmt"

func main() {
	fmt.Println("Valeur ?")
	var input string
	fmt.Scan(&input)
	fmt.Printf("Votre valeur %s", input)
}
