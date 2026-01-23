package main
import {
	"fmt"
	"net/http"
}


type Todo struct {
	UserID int `json:"userid"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed `json:"completed"`
}

func main(){
	fmt.Println("Learning POST request in GO")
    response , error := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if error != nil {
		fmt.Println("Error while getting response", error)
	}
	defer response.Body.Close()
    if response.StatusCode != http.StatusOK {
		fmt.Println("Error while getting response", response.StatusCode)
	}
	fmt.Println("Response status", response.Status)

	var todo Todo
	error = json.NewDecoder(response.body).Decode(&Todo)
	if error != nil {
		fmt.Println("error while getting the data " , error)
	}
	fmt.Println("Todo",todo)
}