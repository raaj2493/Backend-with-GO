
/*
File: functions.go

Purpose:
Covers basic function syntax in Go with minimal notes.
*/

package main

import "fmt"

func main() {

	// Simple function call
	greet()

	// Function with parameters
	sum := add(10, 20)
	fmt.Println("Sum:", sum)

	// Multiple return values
	result, status := divide(10, 2)
	fmt.Println("Result:", result, "Status:", status)

	// Named return values
	fmt.Println("Area:", rectangleArea(5, 4))

	// Variadic function
	fmt.Println("Total:", total(1, 2, 3, 4))

	// Anonymous function
	func(message string) {
		fmt.Println(message)
	}("Hello from anonymous function")

	// Function as a value
	fn := multiply
	fmt.Println("Multiply:", fn(3, 4))

	// Function returning a function
	next := counter()
	fmt.Println(next())
	fmt.Println(next())
}

/* Simple function */
func greet() {
	fmt.Println("Hello, Go!")
}

/* Function with parameters and return value */
func add(a int, b int) int {
	return a + b
}

/* Function with multiple return values */
func divide(a int, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

/* Named return values */
func rectangleArea(length, width int) (area int) {
	area = length * width
	return
}

/* Variadic function */
func total(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

/* Function as value */
func multiply(a, b int) int {
	return a * b
}

/* Function returning function (closure) */
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
