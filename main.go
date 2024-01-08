package main

import "fmt"
import "rsc.io/quote"

func main() {
	fullname := "Akhdan Robbani"
	age := 20
	town := "Nganjuk"
	country := "Indonesia"

	fmt.Println("Hello, World!!!")
	fmt.Println(hello(fullname))
	fmt.Println(age, "years old")
	fmt.Println(address(town, country))
	fmt.Println("quote rsc go :", quote.Go())
}

func hello(name string) string {
	message := fmt.Sprintf("%v in here", name)
	return message
}

func address(town string, country string) string {
	return fmt.Sprintf("lived in %v, %v", town, country)
}
