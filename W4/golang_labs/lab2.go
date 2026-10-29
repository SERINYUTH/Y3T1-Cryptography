package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Enter first integer (x): ")
	fmt.Scan(&x)
	fmt.Print("Enter second integer (y): ")
	fmt.Scan(&y)

	// 1. Both positive (&&)
	bothPositive := (x > 0) && (y > 0)
	fmt.Printf("Are both numbers positive? %t\n", bothPositive)

	// 2. One greater than the other (||)
	oneGreater := (x > y) || (y > x)
	fmt.Printf("Is one greater than the other? %t\n", oneGreater)

	// 3. Not equal (!)
	notEqual := !(x == y)
	fmt.Printf("Are the numbers not equal? %t\n", notEqual)
}
