package main

import "fmt"

func main(){
	firstname, lastname := multiValue()
	fmt.Println(firstname, lastname)
}

func multiValue() (string, string){

	return "akhdan", "robbani"
}