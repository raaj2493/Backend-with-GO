
/*
File: maps.go

Purpose:
Demonstrates maps in Go.
Covers creation, access, update, deletion,
existence check, iteration, and common rules.

Maps are:
- Key-value data structures
- Reference types
- Unordered
*/

package main

import "fmt"

func main() {

	/*
	--------------------------------------------------------
	1. Creating Maps
	--------------------------------------------------------
	*/

	// Using map literal
	ages := map[string]int{
		"Raj":   21,
		"Aman":  22,
		"Riya":  20,
	}
	fmt.Println("Ages:", ages)

	// Using make()
	scores := make(map[string]int)
	scores["Math"] = 90
	scores["Science"] = 85

	fmt.Println("Scores:", scores)

	/*
	--------------------------------------------------------
	2. Accessing Values
	--------------------------------------------------------
	*/

	fmt.Println("Raj's age:", ages["Raj"])

	// Accessing non-existing key returns zero value
	fmt.Println("Unknown age:", ages["Unknown"])

	/*
	--------------------------------------------------------
	3. Checking Key Existence (comma ok idiom)
	--------------------------------------------------------
	*/

	value, exists := ages["Aman"]
	if exists {
		fmt.Println("Aman exists with age:", value)
	} else {
		fmt.Println("Aman does not exist")
	}

	/*
	--------------------------------------------------------
	4. Updating Values
	--------------------------------------------------------
	*/

	ages["Raj"] = 22
	fmt.Println("Updated Ages:", ages)

	/*
	--------------------------------------------------------
	5. Deleting Keys
	--------------------------------------------------------
	*/

	delete(ages, "Riya")
	fmt.Println("After deletion:", ages)

	/*
	--------------------------------------------------------
	6. Iterating Over Maps
	--------------------------------------------------------
	- Order is NOT guaranteed
	*/

	for key, value := range ages {
		fmt.Println("Key:", key, "Value:", value)
	}

	/*
	--------------------------------------------------------
	7. Zero Value of Map
	--------------------------------------------------------
	*/

	var nilMap map[string]int
	fmt.Println("nilMap:", nilMap)

	// Writing to a nil map causes panic
	// nilMap["test"] = 1  // ❌ panic

	/*
	--------------------------------------------------------
	8. Maps are Reference Types
	--------------------------------------------------------
	*/

	original := map[string]int{"a": 1, "b": 2}
	copyMap := original

	copyMap["a"] = 100

	fmt.Println("Original:", original)
	fmt.Println("CopyMap:", copyMap)

	/*
	--------------------------------------------------------
	9. Clearing a Map
	--------------------------------------------------------
	*/

	for key := range original {
		delete(original, key)
	}
	fmt.Println("Cleared map:", original)

	/*
	--------------------------------------------------------
	10. Common Map Rules
	--------------------------------------------------------
	✔ Keys must be comparable types
	✔ Values can be any type
	✔ Maps are unordered
	✔ Zero value is nil
	✔ Use comma-ok to check existence

	--------------------------------------------------------
	End of maps demo
	--------------------------------------------------------
	*/
}
