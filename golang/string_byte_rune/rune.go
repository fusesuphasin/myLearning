package main

import "fmt"

func main2(){
	// []bybe, byte => (0-255)
	str := "ฟิวส์"
	first := str[0]

	fmt.Println(first)
	fmt.Println(string(first))

	str1 := "ฟ"

	fmt.Println([]byte(str1))
	fmt.Println(string([]byte(str1)))

	//rune => int32
	words := "ฟิวส์กหกหก"
	word := []rune(words)

	fmt.Println(word)
	fmt.Println(len(word))
	fmt.Println(string(word))
}