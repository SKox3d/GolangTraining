package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name   string
	Age    int
	Health string
}

func (p Person) getFullName() string {
	return p.Name + ", " + strconv.Itoa(p.Age)
}

type Two struct {
	Person
	Name          string
	LicenceToKill bool
}

func main() {
	a := Person{"Constantin", 26, "Good"}
	fmt.Printf("Name: %s \nAge: %v \nHealth: %s\n", a.Name, a.Age, a.Health)

	b := Two{
		Person: Person{
			Name:   "Ion",
			Age:    25,
			Health: "Good mood",
		},
		Name:          "Vanea",
		LicenceToKill: true,
	}
	fmt.Println(b)

	fmt.Println("Numele complet este :", a.getFullName())
}
