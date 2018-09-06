package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()

	wg.Add(2)
	go first()
	go second()
	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("Script took %s", elapsed)
}

func first() {
	for i := 0; i < 10000; i++ {
		fmt.Println("First: ", i)
		//time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}
func second() {
	for i := 0; i < 10000; i++ {
		fmt.Println("Second: ", i)
		//time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}
