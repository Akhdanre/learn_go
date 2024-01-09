package main

import "fmt"

func main() {
	//array tidak dinamsi sehingga harus di inisiasi jumlah valuenya
	var names [2]string
	names[0] = "AKhdan"
	names[1] = "Robbani"

	var values = [2]int{
		12, 22,
	}

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])


	fmt.Println(values)
}
