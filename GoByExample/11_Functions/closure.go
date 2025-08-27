package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	f := intSeq()
	g := intSeq()

	for range 5 {
		g()
		fmt.Println("f:", f(), ", g:", g())
	}
}
