package main

import (
	"fmt"
	"strings"
)

func main() {
	result1 := strings.Contains("Babelcoder", "bel")
	result2 := strings.Count("Babelcoder bel", "bel")
	result3 := strings.HasPrefix("bel Babelcoder", "bel")
	result4 := strings.HasSuffix("Babelcoder bel", "bel")
	result5 := strings.Join([]string{"Babelcoder", "bel"}, " ")
	result6 := strings.ToUpper("Babelcoder")
	result7 := strings.ToLower("Babelcoder")

fmt.Println(result1)
fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)
}
