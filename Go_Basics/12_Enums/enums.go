// enums.go
package main

import "fmt"

/*
Go does NOT have enums like Java or C++.
Instead, we use:
- a custom type
- constants
This gives type safety and readability.
*/

// ---------- BASIC ENUM ----------

// Status is a custom type based on int
type Status int

// iota auto-increments values (0, 1, 2, ...)
const (
	Pending Status = iota // 0
	Approved              // 1
	Rejected              // 2
)

// ---------- ENUM WITH STRING VALUES ----------

// Role is a custom type based on string
type Role string

const (
	Admin Role = "ADMIN"
	User  Role = "USER"
	Guest Role = "GUEST"
)

// ---------- ENUM WITH START VALUE ----------

// Priority enum starting from 1 instead of 0
type Priority int

const (
	Low Priority = iota + 1 // 1
	Medium                  // 2
	High                    // 3
)

// ---------- ENUM WITH BIT FLAGS (ADVANCED BUT COMMON) ----------

// Permission uses bitwise flags
type Permission int

const (
	Read Permission = 1 << iota // 1 (001)
	Write                       // 2 (010)
	Execute                     // 4 (100)
)

// ---------- MAIN FUNCTION (USAGE EXAMPLES) ----------

func main() {
	var s Status = Approved
	fmt.Println("Status:", s)

	var r Role = Admin
	fmt.Println("Role:", r)

	var p Priority = High
	fmt.Println("Priority:", p)

	// Combine permissions using OR
	var perm Permission = Read | Write
	fmt.Println("Permission:", perm)

	// Check permission using AND
	if perm&Read != 0 {
		fmt.Println("Read permission granted")
	}
}
