package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int `json:"'asd score',wear"`
}

func main() {
	p1 := Person{"Constantin", "Sipitca", 26}
	bs1, _ := json.Marshal(p1)
	fmt.Println(bs1)
	fmt.Printf("%T \n", bs1)
	fmt.Println(string(bs1))

	bs2 := []byte(`{"First":"Constantin","Last":"Harea","asd score":26}`)
	json.Unmarshal(bs2, &p1)
	fmt.Println("-----------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

}
