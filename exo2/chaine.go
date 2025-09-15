package main

import (
	"fmt"
	"strconv"
)

func ConcatLines(args ...interface{}) string {
	var result string
	for i, arg := range args {
		switch v := arg.(type) {
		case string:
			result += v
		case int:
			result += strconv.Itoa(v)
		default:
			result += fmt.Sprintf("%v", v)
		}
		if i < len(args)-1 {
			result += "\n"
		}
	}
	return result
}

func main() {
	fmt.Println("Test 1:")
	fmt.Println(ConcatLines("jouer", 42, "Go", 2025))

	fmt.Println("\nTest 2:")
	fmt.Println(ConcatLines(1, 2, 3))
}
