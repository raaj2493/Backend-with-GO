package main

import "fmt"

func main(){
	age := 19

	if age>19 {
		fmt.Println("Person is an adult")
	} else if age>13 {
		fmt.Println("Person is a teenager")
	} else {
		fmt.Println("Person is a kid")
	}

	// Switch 

	i := 9

	switch i {
		case 1 : 
				fmt.Println("ONE")
		case 2 : 
				fmt.Println("TWO")
		case 3 : 
				fmt.Println("THREE")
		case 4 : 
				fmt.Println("FOUR")
		default :
				fmt.Println("Default")
	}
}
