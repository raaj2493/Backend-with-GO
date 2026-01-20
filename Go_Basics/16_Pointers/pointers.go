package main

import "fmt"

func modifyValueByReference(nums *int){
	*nums = *nums * 10
}

func main () {
	num := 7
	ptr := &num;   
	 /* & is used for pointing to the varible , 
	 ki jo pointer hum use kr rhe ha wo kis varible ka address store krega   
	 */
	 fmt.Println("Pointer is pointing to adress : " , ptr)
	 fmt.Println("Data contained through Pointer : ", *ptr)

	 value := 10
	 modifyValueByReference(&value)
	 fmt.Println("New modified value :", value)
}