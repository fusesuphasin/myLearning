package main

import (
	"fmt"
	"suphasin/shape"
)


func main() {
	var s shape.Shape
	s = shape.Circle{Radius: 10}

	c, ok := s.(shape.Circle)

	fmt.Println(c, ok)
	fmt.Println(c.Area())

	var ss shape.Shape
	ss = shape.Rectangle{Width: 2,Height: 2}

	cc, ok := ss.(shape.Rectangle)
	fmt.Println(cc, ok)
	fmt.Println(cc.Area())

	switch ss.(type) {
	case shape.Circle:
		fmt.Println("Circle")
	case shape.Rectangle:
		fmt.Println("Ragtangle")
	}
}