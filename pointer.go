package main

import "fmt"

type Address struct {
	City , province, Country string
}

func main() {
	//pass by value
	address1 := Address{"Nganjuk", "Jatim", "indonesia"}
	// address2 := address1 //copy value
	

	//tambah & untuk pass by value dengan pointer
	var address2 *Address = &address1 // penambahan pointer dengan bintang serta operator dan 
	// address2 := &address1 //versi simple



	address2.province = "Jawa Timur"

	fmt.Println(address1)//ikut berubah
	fmt.Println(address2)
}