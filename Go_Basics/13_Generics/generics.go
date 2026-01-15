// generics.go
package main

import "fmt"

/*
Generics were introduced in Go 1.18+
They allow writing reusable code
without losing type safety.
*/

// ---------- BASIC GENERIC FUNCTION ----------

// add works with any numeric type
func add[T int | float64](a, b T) T {
	return a + b
}

// ---------- GENERIC FUNCTION WITH ANY ----------

// printValue accepts ANY type
func printValue[T any](value T) {
	fmt.Println(value)
}

// ---------- GENERIC SWAP FUNCTION ----------

// swap swaps values of any type
func swap[T any](a, b T) (T, T) {
	return b, a
}

// ---------- GENERIC CONSTRAINT USING INTERFACE ----------

// Number constraint limits allowed types
type Number interface {
	int | int64 | float64
}

// sum works only with Number types
func sum[T Number](a, b T) T {
	return a + b
}

// ---------- GENERIC STRUCT ----------

// Box can hold any type
type Box[T any] struct {
	Value T
}

// getValue returns the stored value
func (b Box[T]) getValue() T {
	return b.Value
}

// ---------- MAIN FUNCTION (USAGE) ----------

func main() {

	// Generic function usage
	fmt.Println(add(10, 20))
	fmt.Println(add(10.5, 5.5))

	// Any type allowed
	printValue("Hello Go")
	printValue(100)
	printValue(true)

	// Swap values
	a, b := swap(1, 2)
	fmt.Println(a, b)

	x, y := swap("Go", "Lang")
	fmt.Println(x, y)

	// Generic constraint usage
	fmt.Println(sum(5, 10))
	fmt.Println(sum(2.5, 3.5))

	// Generic struct usage
	intBox := Box[int]{Value: 100}
	strBox := Box[string]{Value: "Generics"}

	fmt.Println(intBox.getValue())
	fmt.Println(strBox.getValue())
}
