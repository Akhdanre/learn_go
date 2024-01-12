package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string){
	fmt.Println("Hello", name, "my name is", customer.Name)
}
func main(){
	var akhdan Customer
	fmt.Println(akhdan)

	akhdan.Name = "Akhdan Robbani"
	akhdan.Address = "nganjuk"
	akhdan.Age = 20

	fmt.Println(akhdan)
	fmt.Println(akhdan.Name)
	fmt.Println(akhdan.Address)
	fmt.Println(akhdan.Age)


	//cara lain
	ouken := Customer{
		Name: "oukenzeumasio",
		Address: "Nganjuk",
		Age: 20,
	}
	fmt.Println(ouken)

	akeon := Customer{"akeoneuefo", "nganjuk", 20}
	fmt.Println(akeon)

	akeon.sayHello("akhdan")
}