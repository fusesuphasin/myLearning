package main

import "fmt"

//if panic is occer โปรแกรมจะหยุดทำงานทะนที
//ใช้กับ DB
func main(){
	fmt.Println(1)
	fmt.Println(2)
	panic("Fail")
	fmt.Println(4)
}