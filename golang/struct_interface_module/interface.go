package main

import "fmt"

type generator interface { //สนใจแค่พฤษติกรรม / เมดทอด
	generate()
}

type pdf struct {
	content string
}

func (p pdf) generate() { // pdf เป็น generator ด้วย
	fmt.Println("Genereting...")
}

func main() {
	var gen generator
	gen = pdf{content: "My PDF"}

	gen.generate()

	//emty interface
	var s interface{}
	s = "String"
	fmt.Println(s)

	result := []interface{}{"A",false,10}
	fmt.Println(result)

	
}