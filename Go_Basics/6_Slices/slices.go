package main

import "fmt"

/*
File: slices.go

Purpose:
Demonstrates slices in Go.
Covers creation, indexing, length vs capacity,
append, slicing, copying, and common rules.

Slices are:
- Dynamic
- Reference-based
- Built on top of arrays
*/

func main() {

	/*
	--------------------------------------------------------
	1. Creating Slices
	--------------------------------------------------------
	*/

	// Using slice literal
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// Using make()
	// make(type, length, capacity)
	scores := make([]int, 3, 5)
	fmt.Println("Scores:", scores)

	/*
	--------------------------------------------------------
	2. Length vs Capacity
	--------------------------------------------------------
	- len(): number of elements in slice
	- cap(): underlying array capacity
	*/

	fmt.Println("len:", len(scores))
	fmt.Println("cap:", cap(scores))

	/*
	--------------------------------------------------------
	3. Indexing & Updating
	--------------------------------------------------------
	*/

	numbers[0] = 100
	fmt.Println("Updated Numbers:", numbers)

	/*
	--------------------------------------------------------
	4. Appending to Slices
	--------------------------------------------------------
	- append returns a NEW slice
	- capacity may grow automatically
	*/

	numbers = append(numbers, 6)
	numbers = append(numbers, 7, 8)

	fmt.Println("After append:", numbers)
	fmt.Println("len:", len(numbers), "cap:", cap(numbers))

	/*
	--------------------------------------------------------
	5. Slicing a Slice
	--------------------------------------------------------
	Syntax:
	slice[low : high]

	- low is inclusive
	- high is exclusive
	*/

	sub := numbers[1:4]
	fmt.Println("Sub slice:", sub)

	/*
	--------------------------------------------------------
	6. Slices are Reference Types
	--------------------------------------------------------
	Changes affect the underlying array
	*/

	sub[0] = 999
	fmt.Println("Numbers after sub change:", numbers)
	fmt.Println("Sub slice:", sub)

	/*
	--------------------------------------------------------
	7. Copying Slices
	--------------------------------------------------------
	Use copy() to create independent slices
	*/

	original := []int{1, 2, 3}
	duplicate := make([]int, len(original))

	copy(duplicate, original)
	duplicate[0] = 100

	fmt.Println("Original:", original)
	fmt.Println("Duplicate:", duplicate)

	/*
	--------------------------------------------------------
	8. Nil vs Empty Slice
	--------------------------------------------------------
	*/

	var nilSlice []int
	emptySlice := []int{}

	fmt.Println("nilSlice:", nilSlice, "len:", len(nilSlice))
	fmt.Println("emptySlice:", emptySlice, "len:", len(emptySlice))

	/*
	--------------------------------------------------------
	9. Iterating Over Slices
	--------------------------------------------------------
	*/

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index:", i, "Value:", numbers[i])
	}

	for index, value := range numbers {
		fmt.Println("range ->", index, value)
	}

	/*
	--------------------------------------------------------
	10. Common Slice Rules
	--------------------------------------------------------
	✔ Zero value of slice is nil
	✔ append may reallocate memory
	✔ Always assign result of append
	✔ copy creates independent slice

	--------------------------------------------------------
	End of slices demo
	--------------------------------------------------------
	*/
}

