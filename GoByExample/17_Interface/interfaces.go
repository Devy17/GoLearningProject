package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}

func detectCircle(g geometry) {
	/*
		Go는 If, For, Switch에서 선언과 조건을 같이 사용할 수 있음
		ex)
		if 초기화문; 조건문 {
			// 실행문
		}
	*/
	if c, ok := g.(circle); ok {
		fmt.Println("radius: ", c.radius)
	}

}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 7}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
