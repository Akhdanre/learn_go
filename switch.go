package main

import "fmt"

func main(){
	name := "dfd"

	switch name {
	case "Akhdan":
		fmt.Println("hallo Akhdan")
	case "Akeon":
		fmt.Println("hallo Akeon")
	default :
		fmt.Println("halo siapa kamu")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("panjang nama sesuai")
	case false:
		fmt.Println("panjang tidak sesuai")	
	}


	name = "Akeon"
	length := len(name)
	switch {
	case length > 5: 
		fmt.Println("yeay sudah benar")
	case length <= 5:
		fmt.Println("kurang sesuai")
	}
}