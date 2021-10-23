package main

import "fmt"

func g(){
	
}

func h() {
	
	defer func() {
		if r := recover(); r != nil{ // r != nil -> เกิด panic 
			fmt.Println(r, "Recover") 
		}
	}()
	panic("False")
	g()
}

func main(){
	h()
	fmt.Println("2")
}