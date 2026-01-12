
/*
File: closures.go

Purpose:
Shows what a closure is in the simplest way.
*/

package main

import "fmt"

func main() {

	// Closure keeps its own state
	counter := createCounter()
	fmt.Println(counter())
	fmt.Println(counter())

	// New closure = new state
	anotherCounter := createCounter()
	fmt.Println(anotherCounter())

	// Closure using outer variable
	double := multiplyBy(2)
	fmt.Println(double(5))
}

/* Closure that remembers count */
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

/* Closure capturing a value */
func multiplyBy(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
