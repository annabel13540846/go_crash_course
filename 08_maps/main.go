package main

import "fmt"

func main() {

	// Define map
	//emails := make(map[string]string)

	// Assign key values
	// emails["Bob"] = "bob@gmail.com"
	// emails["Chloe"] = "chloe@gmail.com"
	// emails["Sasha"] = "sasha@gmail.com"

	// Declare map and add key values

	emails := map[string]string{"Bob": "bob@gmail.com", "Chloe": "chloe@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	delete(emails, "Bob")
	fmt.Println(emails)
}
