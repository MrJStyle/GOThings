package main

import (
	"basic/struct/entities"
	"fmt"
)

func main() {
	// cannot use unexported variable to initialize
	p := entities.Person1{Name: "boy"}
	fmt.Println(p)

	// cannot directly use unexported member 'person2' to initialize, but can set exported member
	a := entities.Admin{Level: "high"}
	a.Name = "Cindy"
	a.Email = "xxxxx@gmail.com"
	fmt.Println(a)
}
