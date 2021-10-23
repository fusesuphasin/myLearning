package main

import "fmt"

type users struct {
	name string
	age  uint
}

type article struct{
	title string
	excerpt string
	body string
	users users
}

func main() {
	Me := users{name: "Suphasin",age: 22}
	fmt.Println(Me)
	fmt.Println(Me.name)

	a := article{
		title: "Title",
		excerpt: "Excerpt",
		body: "Body",
		users: users{name: "Suphasin",age: 22},
	}

	fmt.Println(a)
}