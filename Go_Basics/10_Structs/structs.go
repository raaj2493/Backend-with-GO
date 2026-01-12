package main

import "fmt"

func main() {

	// Creating struct using named fields
	user1 := User{
		Name: "Raj",
		Age:  21,
		City: "Delhi",
	}

	fmt.Println(user1)

	// Accessing fields
	fmt.Println(user1.Name, user1.Age)

	// Updating fields
	user1.Age = 22
	fmt.Println("Updated age:", user1.Age)

	// Creating struct without field names
	user2 := User{"Aman", 20, "Noida"}
	fmt.Println(user2)

	// Pointer to struct
	userPtr := &user1
	userPtr.City = "Gurgaon"
	fmt.Println("Updated city:", user1.City)

	// Anonymous struct
	temp := struct {
		ID   int
		Name string
	}{
		ID:   1,
		Name: "Temp User",
	}
	fmt.Println(temp)
}

/* Simple struct definition */
type User struct {
	Name string
	Age  int
	City string
}
