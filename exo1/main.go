package main

import "fmt"

func Range(min, max int) []int {
	if min >= max {
		return nil
	}

	var result []int
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}

func main() {
	fmt.Println(Range(5, 10))
	fmt.Println(Range(10, 5))

}
