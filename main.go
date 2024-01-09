package main

import "fmt"
import "rsc.io/quote"

func main() {
	// introduction();

}

func introduction(){
	const fullname string = "Akhdan Robbani" //value yang tetap atau tidak bisa berubah valuenya
	// const (
		// firstname = "Akhdan"
		// lastname = "Robbani"
		// )
	
	// deklarasi multi variablejj
	var (
		firstname = "Akhdan"
		lastname = "Robbani"
	)
	
	var age = 20 // tanpa deklarasi tipe data karena sudah memiliki default value


	var town string //
	town = "nganjuk"
	town = "Nganjuk"

	country := "Indonesia"

	fmt.Println("Hello, World!!!")
	fmt.Println(hello(fullname))
	fmt.Println("firstname :", firstname)
	fmt.Println("lastname :", lastname)
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



