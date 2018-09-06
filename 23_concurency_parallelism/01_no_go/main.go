package main

import (
	"fmt"
	"log"
	"time"
)

var firstNum int64
var secondNum int64

func main() {
	start := time.Now()

	first()
	second()

	fmt.Println("First: ", firstNum)
	fmt.Println("Second: ", secondNum)

	elapsed := time.Since(start)
	log.Printf("Script took %s", elapsed)
}

func first() {
	firstNum++
	time.Sleep(3 * time.Second)
}
func second() {
	secondNum++
	time.Sleep(2 * time.Second)
}
