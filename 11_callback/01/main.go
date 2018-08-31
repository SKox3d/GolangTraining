package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func main() {

	visit([]int{3, 2, 1, 4}, mtd)
}

func mtd(n int) {
	fmt.Println(n)
}
