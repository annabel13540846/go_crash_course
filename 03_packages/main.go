package main

import (
	"fmt"
	"math"

	"github.com/annabel13540846/go_crash_course/03_packages/strutil"
)

func main() {

	fmt.Println("Round down 2.7 -", math.Floor(2.7))
	fmt.Println("Round up 2.7 -", math.Ceil(2.7))
	fmt.Println(strutil.Reverse("olleh"))

}
