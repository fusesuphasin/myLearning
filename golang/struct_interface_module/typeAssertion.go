package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct{
	radius float64
}

type rectangle struct{
	width float64
	height float64
}
 
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return 2 * (r.width + r.height)
}

func main(){
	var s shape
	s = circle{radius: 10}
	
	c, ok := s.(circle)

	fmt.Println(c, ok)
	fmt.Println(c.area())

	var ss shape
	ss = rectangle{width: 2, height: 2}

	cc, ok := ss.(rectangle)
	fmt.Println(cc, ok)
	fmt.Println(cc.area())

	switch ss.(type) {
	case circle:
		fmt.Println("Circle")
	case rectangle:
		fmt.Println("Ragtangle")
	}
}