package main


import (
	"fmt"
	"net/url"
)

func main (){
	fmt.Println("Learning Web in GO")
	myURL := "https://www.youtube.com/watch?v=FpOcxFBzCow"
	fmt.Printf("Type of URL: %T\n", myURL)

	parsedURL , err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Cannot parse URL",err)
	}

	fmt.Printf("Parsed URL: %T\n", parsedURL)

	fmt.Println("Scheme:",parsedURL.Scheme)
	fmt.Println("Host:",parsedURL.Host)
	fmt.Println("Path:",parsedURL.Path)
	fmt.Println("Query:",parsedURL.RawQuery)

}