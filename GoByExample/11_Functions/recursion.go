package main

import "fmt"

func fact(n int) int {
	if n == 1 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// 익명 함수도 가능.. 은 한데 어차피 메모리 올려놔야 함
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
