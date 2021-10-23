package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printAndSleep(num int) {
	defer wg.Done() 

	time.Sleep(1 * time.Second)
	fmt.Println(num)
}

func main() {
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go printAndSleep(i)
	}

	wg.Wait()
}