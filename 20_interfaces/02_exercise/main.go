package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []string{"asd", "dsa", "das", "ads"}

	fmt.Println(n)
	sort.Sort(sort.StringSlice(n))
	fmt.Println(n)

}
