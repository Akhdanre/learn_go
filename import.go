package main

import "learn_go/helper"
import "fmt"

func main(){
	result := helper.SayHello("Akhdan")
	fmt.Println(result)

	fmt.Println(helper.Application)
	//tidak bisa diakses
	// fmt.Println(helper.version)
	// fmt.Println(helper.sayGoodBye())
}