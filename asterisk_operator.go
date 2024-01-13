package main

import "fmt"

type Address struct {
	City , province, Country string
}

func main() {
	var address1 Address = Address{"Nganjuk", "Jatim", "Indonesia"}
	var address2 *Address = &address1

	address2.City= "Kediri"

	fmt.Println(address1)
	fmt.Println(address2)
	
	// address2 = &Address{"Jombang", "Jawa timur", "Indonesia"}

	//mebgubah value address2 dan semua yang mengacu
	*address2 = Address{"Jombang", "Jawa timur", "Indonesia"}
	
	fmt.Println(address1)
	fmt.Println(address2)
}