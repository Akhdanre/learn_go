package main

import "fmt"

type Address struct {
	City, Province, Country string
}


func ChangeCountry(address *Address){
	address.Country = "indonesia"
}


func main(){
	address  := Address{}
	ChangeCountry(&address)

	fmt.Println(address)
}