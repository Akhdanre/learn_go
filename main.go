package main

import "fmt"
import "rsc.io/quote"

func main() {
	var fullname string = "Akhdan Robbani"
	var age int8 = 20 // age := 20

	var town string 
	town = "nganjuk"
	town = "Nganjuk"

	country := "Indonesia"

	fmt.Println("Hello, World!!!")
	fmt.Println(hello(fullname))
	fmt.Println(age, "years old")
	fmt.Println(address(town, country))
	fmt.Println("quote rsc go :", quote.Go())
	fmt.Println(len("mengambil panjang string dengan len"))
}

func hello(name string) string {
	message := fmt.Sprintf("%v in here", name)
	return message
}

func address(town string, country string) string {
	return fmt.Sprintf("lived in %v, %v", town, country)
}
