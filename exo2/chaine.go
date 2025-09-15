package main

import (
	"fmt"
	"strings"
)

func assemblerEnChaine(args ...interface{}) string {
	var lignes []string
	for _, arg := range args {
		lignes = append(lignes, fmt.Sprintf("%v", arg))
	}
	return strings.Join(lignes, "\n")
}

func main() {
	fmt.Println("Test 1:")
	fmt.Println(assemblerEnChaine("Bonjour", 42, "Aurevoir"))

	fmt.Println("\nTest 2:")
	fmt.Println(assemblerEnChaine(1, 2, 3, 4, 5))

}
