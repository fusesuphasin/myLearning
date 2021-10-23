package main

import "fmt"

type Books struct{
	title string
	auther string
	subtible string
	price float64 
}

func main5(){
	var Book1 Books
	Book1.title = "MATH"
	Book1.auther = "FUSE"
	Book1.subtible = "MY WORLD"
	Book1.price = 0.0

	fmt.Println(Book1)
	fmt.Println(Book1.auther)

	Book2:=Books{title: "ART", auther: "WHO", subtible: "IDK", price: 1}
	fmt.Println(Book2)
	fmt.Println(Book2.auther)
}