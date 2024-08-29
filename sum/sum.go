package main

import "fmt"

func main() {
	var i int
	var j int

	fmt.Println("Enter a number:")
	fmt.Scan(&i)
	fmt.Println("Enter another number:")
	fmt.Scan(&j)

	fmt.Println("Sum of", i, "and", j, "is", i+j)
}
