package main

import (
	"fmt"
	"strconv"
)

//Biblioteca strconv godoc.org
func main() {
	s := "2147483647" // biggest int32
	i64, _ := strconv.ParseInt(s, 10, 32)

	i := int64(i64)
	fmt.Println(i)
}
