package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	fmt.Printf("%T", callback)
	return xs
}

func main() {
	arr := []int{2, 3, 4, 1}
	verify := func(n int) bool {
		return n > 1
	}

	xs := filter(arr, verify)
	fmt.Println(xs)

}
