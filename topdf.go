package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter Your Number: ")
	var first int
	fmt.Scanln(&first)

	if first%2 == 0 {
		fmt.Print(first, " is an even number consisting of ")
		for i := 0; i <= first; i = i + 2 {
			fmt.Print(" ", i)
		}

	} else {
		fmt.Print(first, " is an odd number consisting of ")
		for i := 1; i <= first; i = i + 2 {
			fmt.Print(" ", i)
		}
	}
}
