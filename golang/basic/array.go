package main

import "fmt"

func main1(){
	//array ขนาดตายตัว
	/* var array[5] int
	array[2] = 1
	fmt.Println(array)
	fmt.Println(array[2])

	x:=[10]float64{0,5,10,15,20,25}
	fmt.Println(x)
	var total float64
	for _ , number:= range x {
		total += number
	}
	fmt.Println(total)
	fmt.Println(total/float64(len(x))) */

	//slice ขนาดยิดหยุ่นได้
	/* slice1:=make([]int,10)
	fmt.Println(slice1)
    
	var a []int =[]int{1,2,3}
	fmt.Println(a)

	var a1 [5]int =[5]int{1,2,3,4,5}
	fmt.Println(a1)

	slice2:=[]int{1,2,3}
	fmt.Println(slice2)

	slice3:=append(slice2,4,5)
	fmt.Println(slice3)

	copy(slice1,slice3)
	fmt.Println(slice1) */

	//map
	y:=make(map[string] string)
	y["TH"]="THAILAND"
	y["JP"]="JAPAN"
	y["EN"]="ENGLAND"
	fmt.Println(y["EN"])

	var mp map[string]int = map[string]int{
		"apple": 5,
		"pear": 3,
	}
	val, ok := mp["apple"]
	fmt.Println(val , ok)

	z:=map[string] string{
		"TH":"THAILAND",
		"JP":"JAPAN",
	}
	fmt.Println(z)

	/* 
	//for-range
	name := []string{"1","2","3"}
	for k, v := range name {
		k = index
		v = value
		but use name = object etc "name" : "su"
		k = name
		v = su
	}
	*/
}