/*
File: variables.go

Purpose:
A single, clean reference file covering everything about variables in Go.
Includes declarations, types, constants, scope, pointers, and rules.

How to use:
- Read section by section
- Run and experiment
- Keep this as a long-term reference
*/

package main

import "fmt"

/*
------------------------------------------------------------
Global Declarations
------------------------------------------------------------
- Global variables must use `var`
- Short declaration `:=` is NOT allowed at package level
*/

var globalCount int = 100
var globalMessage = "Hello from global scope"

const globalConst = "I am constant"

func main() {

	/*
	--------------------------------------------------------
	1. Basic Variable Declaration
	--------------------------------------------------------
	Explicit type declaration using `var`
	*/

	var age int = 21
	var name string = "Raj"
	var isStudent bool = true

	fmt.Println(age, name, isStudent)

	/*
	--------------------------------------------------------
	2. Type Inference
	--------------------------------------------------------
	Go infers the type from the assigned value
	*/

	var city = "Delhi"
	var score = 95.5

	fmt.Println(city, score)

	/*
	--------------------------------------------------------
	3. Zero Values
	--------------------------------------------------------
	Uninitialized variables receive default zero values
	*/

	var a int
	var b float64
	var c bool
	var d string
	var e []int
	var f *int

	fmt.Println(a, b, c, d, e, f)

	/*
	--------------------------------------------------------
	4. Short Variable Declaration (:=)
	--------------------------------------------------------
	- Most commonly used form
	- Allowed only inside functions
	*/

	language := "Go"
	version := 1.22

	fmt.Println(language, version)

	/*
	--------------------------------------------------------
	5. Multiple Variable Declaration
	--------------------------------------------------------
	*/

	var x, y, z int = 1, 2, 3
	fmt.Println(x, y, z)

	var username, loggedIn = "admin", true
	fmt.Println(username, loggedIn)

	/*
	--------------------------------------------------------
	6. Variable Block
	--------------------------------------------------------
	Used to group related variables for cleaner code
	*/

	var (
		productName  = "Laptop"
		productPrice = 75000
		inStock      = true
	)

	fmt.Println(productName, productPrice, inStock)

	/*
	--------------------------------------------------------
	7. Constants
	--------------------------------------------------------
	- Immutable
	- Must be initialized
	- Evaluated at compile time
	*/

	const pi = 3.14159
	const appName string = "Go Variables"

	fmt.Println(pi, appName)

	/*
	--------------------------------------------------------
	8. Built-in Data Types
	--------------------------------------------------------
	*/

	var i int = 10
	var i8 int8 = 127
	var f32 float32 = 3.14
	var f64 float64 = 3.141592
	var ch byte = 'A'
	var r rune = '₹'
	var s string = "Go"

	fmt.Println(i, i8, f32, f64, ch, r, s)

	/*
	--------------------------------------------------------
	9. Type Conversion
	--------------------------------------------------------
	Go does NOT allow implicit type conversion
	*/

	var num int = 10
	var decimal float64 = float64(num)

	fmt.Println(decimal)

	/*
	--------------------------------------------------------
	10. Scope & Shadowing
	--------------------------------------------------------
	Local variables can shadow outer or global variables
	*/

	globalCount := 50
	fmt.Println("Shadowed globalCount:", globalCount)

	printGlobal()

	xVal := 100
	{
		xVal := 200
		fmt.Println("Inner xVal:", xVal)
	}
	fmt.Println("Outer xVal:", xVal)

	/*
	--------------------------------------------------------
	11. Pointers
	--------------------------------------------------------
	- Store memory addresses
	- Use `&` to get address
	- Use `*` to dereference
	*/

	num1 := 10
	ptr := &num1

	fmt.Println("Value:", num1)
	fmt.Println("Address:", ptr)
	fmt.Println("Via pointer:", *ptr)

	*ptr = 20
	fmt.Println("Updated Value:", num1)

	/*
	--------------------------------------------------------
	12. Unused Variables Rule
	--------------------------------------------------------
	Go enforces usage of all declared variables
	*/

	value, _ := 10, 20
	fmt.Println(value)

	/*
	--------------------------------------------------------
	End of main
	--------------------------------------------------------
	*/
}

/*
------------------------------------------------------------
Helper Function
------------------------------------------------------------
Demonstrates access to global variables
*/

func printGlobal() {
	fmt.Println("Global Count:", globalCount)
	fmt.Println("Global Message:", globalMessage)
	fmt.Println("Global Const:", globalConst)
}
