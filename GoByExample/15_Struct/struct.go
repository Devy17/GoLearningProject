package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 얕은 복사
	sp := s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println("원본 : ", s.age)

	// 깊은 참조 (내부적으로 (*sp2).age로 안써도 되도록 변환)
	sp2 := &s
	fmt.Println(sp2.age)
	sp2.age = 52
	fmt.Println(sp2.age)
	fmt.Println("원본 : ", s.age)

	dog := struct {
		name   string
		isGood bool
	}{
		name:   "Sam",
		isGood: true,
	}
	fmt.Println(dog)
}
