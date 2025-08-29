package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("key:", k, "value:", v)
	}

	// 키만 가지고도 돌릴 수 있음..
	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "가나다" {
		fmt.Println(i, c)
	}
}
