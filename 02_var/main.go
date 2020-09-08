package main

import "fmt"

func main() {
	// Using var
	//var name = "Annabel"
	// Shorthand
	//name := "Brad"
	var age int32 = 20
	const isCool = true
	// size := 1.3
	// var size float32 = 2.3
	name, email := "Brad", "brad@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", email)

}
