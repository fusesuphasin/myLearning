package main

import (
	"fmt"
	"strconv"
)

func main1(){
	// _ คือ ไม่รับ error ถ้ารับใส่ตัวแปรเข้าไป err
	a, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println(a)

	e, _ := strconv.ParseInt("11101",2, 64)
	fmt.Println(e)

	i, _ := strconv.ParseUint("232", 10, 64)
	fmt.Println(i)

	o, _ := strconv.Atoi("65")
	fmt.Println(o)

	w, _ := strconv.Itoa(2)
	fmt.Println(w)
}
