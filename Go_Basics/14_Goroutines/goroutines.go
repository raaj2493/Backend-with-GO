package main

import (
	"fmt"
	"time"
)

func main (){
   for i := 0; i <= 10; i++ {
	go task(i)
   }

   time.Sleep(time.Second*2)
}

func task (id int ){
  fmt.Println("Doing Task no -", id)
}