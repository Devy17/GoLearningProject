package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 1; j < 10; j++ {
		fmt.Println(j)
	}

	// 이거 go 1.22 버전 이상부터 사용 가능
	for k := range 3 {
		fmt.Println(k)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}
