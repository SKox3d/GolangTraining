package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	First string
	Last  string
	Age   int `json:"'asd score',wear"`
}

func main() {
	p1 := Person{"Constantin", "Sipitca", 26}
	json.NewEncoder(os.Stdout).Encode(p1)

	rdr := strings.NewReader(`{"First":"Doinita","Last":"Droganenco","Age":26}`)
	json.NewDecoder(rdr).Decode(&p1)
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
