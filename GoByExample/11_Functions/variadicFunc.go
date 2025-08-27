package main

import "fmt"

func sum(nums ...int) int {
	fmt.Print(nums, ", ")
	total := 0
	// for ... range 문은 리턴이 2개
	// for index, value := range arr { ... }
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
	return total
}

func main() {
	sum(1, 2, 3, 4)
	sum(1, 2, 3, 4, 5)

	l := []int{1, 2, 3, 4}
	sum(l...)
}
