package main

import "fmt"

func main3(){
	//routine ใช้แทรกแซงเหตุการณ์ ไม่ต้องทำตามลำดับขั้นตอน
	go f(0)
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
	

}

func f(n int){
	for i:=0; i<10; i++{
		fmt.Println(n, " : ",i)
	}
}