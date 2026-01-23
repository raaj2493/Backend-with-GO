package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func main() {
	webrequest()
}

func webrequest() {
	fmt.Println("Learning Web Services")
	response , err := http.Get("https://jsonplaceholder.typicode.com/todos/")
	if err != nil {
		fmt.Println("Error while getting response", err)
	}
	defer response.Body.Close()
	fmt.Println("Type of response %T\n" , response)

	// Read the response body 
   data , err := ioutil.ReadAll(response.Body)
   if err != nil {
   	fmt.Println("Error while reading response", err)
   }
   fmt.Println("Response data", string(data))
}