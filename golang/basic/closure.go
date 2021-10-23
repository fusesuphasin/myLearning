package main

import "fmt"

func main2() {

	//closure
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(2, 3))

	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	//pointer
	y := 10
	fmt.Println("value is ", y)
	fmt.Printf("Address x valiable %x \n", &x)
	fmt.Printf("Address y valiable %x \n", &y)

	z := 10
	var p *int
	p = &z //point to valiable z is
	fmt.Printf("Address p = %x \n", p)

	name := "FUSE"
	var pName *string
	pName = &name
	fmt.Printf("Address name = %x \n", pName)

	*pName = "SUPHASIN"
	fmt.Printf("name = %s \n", name)

}
