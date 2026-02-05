/*
File: constants.go

Purpose:
Demonstrates how constants work in Go.
Covers basic constants, typed vs untyped constants,
constant blocks, and important rules.

Notes:
- Constants are immutable
- Constants are evaluated at compile time
- Constants cannot be declared using :=

*/

package main

import "fmt"

/*
------------------------------------------------------------
Global Constants
------------------------------------------------------------
- Can be declared at package level
- Must be initialized at declaration
*/

const appName = "Go Learning App"
const appVersion = "1.0.0"

/*
------------------------------------------------------------
Constant Block
------------------------------------------------------------
Used to group related constants together
*/

const (
	maxUsers      = 100
	defaultPort   = 8080
	enableLogging = true
)

func demonstrateConstants() {

	/*
	--------------------------------------------------------
	1. Typed vs Untyped Constants
	--------------------------------------------------------
	*/

	const untypedPi = 3.14          // type decided at usage
	const typedPi float64 = 3.14159 // fixed type

	fmt.Println(untypedPi, typedPi)

	/*
	--------------------------------------------------------
	2. Constants with Expressions
	--------------------------------------------------------
	*/

	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	fmt.Println("Seconds in an hour:", secondsInHour)

	/*
	--------------------------------------------------------
	3. Numeric Constant Flexibility
	--------------------------------------------------------
	Untyped constants can adapt to required type
	*/

	const value = 10

	var i int = value
	var f float64 = value

	fmt.Println(i, f)

	/*
	--------------------------------------------------------
	4. What Constants Cannot Do
	--------------------------------------------------------
	These are NOT allowed (examples commented)
	*/

	// const now = time.Now()        // ❌ runtime value
	// const x := 10                // ❌ := not allowed
	// const y                     // ❌ must be initialized

	/*
	--------------------------------------------------------
	End of constants demo
	--------------------------------------------------------
	*/
}

func main() {
	demonstrateConstants()
}
