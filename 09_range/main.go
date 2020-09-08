package main

import "fmt"

func main() {

	ids := []int{33, 56, 67, 87, 95, 2, 10}
	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with maps
	emails := map[string]string{"Bob": "bob@gmail.com", "Chloe": "chloe@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
