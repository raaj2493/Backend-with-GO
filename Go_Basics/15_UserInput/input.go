package main

import (
	"bufio"
	"fmt"
	"os"
)

func main (){
	var name string;
	fmt.Scan(&name)
	fmt.Print("hello " , name)

	/*
	 the problem with scan method is that it is only valid till white spaces 
	 means we cannot input whole lines in it. we can only input single words
	*/

	// to fix this problem we have bufio package

	reader := bufio.NewReader(os.Stdin)
	name1 , _ := reader.ReadString('\n')     // "\n" means it will read till next line
	fmt.Println("hello" , name1)
}