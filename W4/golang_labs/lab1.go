package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	//Task 1
	var a, b int
	fmt.Print("Enter first integer (a): ")
	fmt.Scan(&a)
	fmt.Print("Enter second integer (b): ")
	fmt.Scan(&b)

	// Using assignment operators
	c := a
	fmt.Printf("c = a   -> c is now: %d\n", c)

	c += b
	fmt.Printf("c += b  -> c is now: %d\n", c)

	c -= b
	fmt.Printf("c -= b  -> c is now: %d\n", c)

	c *= b
	fmt.Printf("c *= b  -> c is now: %d\n", c)

	c /= b
	fmt.Printf("c /= b  -> c is now: %d\n", c)

	c %= b
	fmt.Printf("c %%= b -> c is now: %d\n", c)
}
