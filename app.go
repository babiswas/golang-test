package main

import (
	"Appp/Person"
	"fmt"
)

func main() {
	p := Person.Person{}
	p.Set_name("Bapan")
	p.Set_id(21)
	fmt.Println(p.Display())
	fmt.Println(p.Work())
	fmt.Println(p.Salary())
}
