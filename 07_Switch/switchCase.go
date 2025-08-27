package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(input interface{}) {
		switch k := input.(type) {
		case bool:
			fmt.Println("It is boolean")
		case int:
			fmt.Println("It is an int")
		default:
			fmt.Printf("I don't know about type %T \n", k)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
