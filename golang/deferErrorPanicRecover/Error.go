package main

import (
	"errors"
	"fmt"
)

func findIndex(s []int, num int) (int, error) {
	for i, n := range s {
		if n == num {
			return i, nil
		}
	}

	return 0, errors.New("Number not found")
}

func main() {
	i, err := findIndex([]int{1, 2, 3}, 2)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(i)
	}
}