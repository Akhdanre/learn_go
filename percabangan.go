package main

import "fmt"

func main() {
	name := "akhdan"

	if name == "Akhdan" {
		fmt.Println("yeay, it's me %v", name)

	} else if name == "akhdan" {
		fmt.Println("hai, akhdan in lowercase all")

	} else {
		fmt.Println("who are you")
	}


	if length := len(name); length >5 {
		fmt.Println("panjang nama sudah sesuai")
	}else {
		fmt.Println("panjang nama tidak sesuai kurang dari 5")
	}
}
