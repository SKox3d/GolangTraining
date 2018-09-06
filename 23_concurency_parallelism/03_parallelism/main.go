package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var fg sync.WaitGroup

var firstNum int64
var secondNum int64

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	start := time.Now()

	wg.Add(2)

	go first()
	go second()

	wg.Wait()

	fmt.Println("First: ", firstNum)
	fmt.Println("Second: ", secondNum)

	elapsed := time.Since(start)
	log.Printf("Script took %s", elapsed)
}

func first() {
	firstNum++
	fg.Add(2)

	go wait()
	go wait()

	fg.Wait()

	wg.Done()
}
func wait() {
	time.Sleep(3 * time.Second)
	fg.Done()
}
func second() {
	secondNum++
	time.Sleep(2 * time.Second)
	wg.Done()
}
