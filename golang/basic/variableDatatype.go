package main

import "fmt"

//SCOPE
//GLOBAL VARIABLE
var gVariable int = 500


func main() {
	//string
	var name string = "Suphasin Yosang"
	var nickname string
	nickname = "Fuse"

	fmt.Println("My name is" + name + "\nMy nickname is " + nickname)

	//int 
	var age int
	age = 20
	// Assignment Expression & Implicit vs Explicit
	// ถ้ากำหนดค่าเริ่มต้นให้ตัวแปรแต่ได้จะใช้ age := 21

	fmt.Printf("%T", age)
	fmt.Printf("%T", age)
	fmt.Printf("%T", age)

	//boolean
	var isEqual bool
	isEqual = true
	fmt.Println(isEqual)
	fmt.Println(name[1:4])

	//GLOBAL VARIABLE
	fmt.Println(gVariable)
}