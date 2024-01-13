package main

import "fmt"

type Address struct {
	City , province, Country string
}

func main() {
	address1 := Address{"Nganjuk", "Jatim", "indonesia"}
	address2 := address1

	address2.City = "Kediri"

	fmt.Println(address1)
	fmt.Println(address2)
}