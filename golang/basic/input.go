package main

import "fmt"

func main4(){
	fmt.Print("Input your Number: ")
	var number int
	fmt.Scanf("%f", &number)
	fmt.Println(number)

	//if-esle
	if( number >= 0 ){
		fmt.Println("Positive Number")
	}else{
		fmt.Println("Negative Number")
	}

	//switch-case
	switch number {
	case 0:fmt.Println("Zero")
	case 1:fmt.Println("One")
	case 2:fmt.Println("Two")
	case 3:fmt.Println("Three")
	default:fmt.Println("WHAT Number")
	}

	//go has only for loop
	for i:=0; i < 6; i++ {
		fmt.Printf("%d", i)
	}

	i:=1
	for i<6 {
		fmt.Printf("%d", i)
		i++
	}

	myName()
	yourName("Oni")
	result := countDigit(6,2)
	fmt.Println(result)

	//Variadic Function
	summation(1,3,54,6,7,8,2)

	//Recursive Function
	fmt.Println(factorail(7))
}

func myName(){
	fmt.Println()
	fmt.Println("panthom")
}

func yourName(name string){
	fmt.Println(name)
}

func countDigit(num1 int, num2 int) /*ค่าที่รีเทิร์นใส่ตามหลัง*/ int{
	result := num1+num2
	return result 
}

func summation(args...int){
	var total int
	// ตำแหน่งเริ่มต้นที่ 0 -> _
	for _, n:=range args{
		total+=n
	}
	fmt.Println(total)
}

func factorail(num int) int{
	if num==0{
		return 1;
	}
	return num * factorail(num-1)
}