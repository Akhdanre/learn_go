package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married(){
	man.Name = "Mr. "+ man.Name
}

func main(){
	Akhdan := Man{"Akhdan"}

	Akhdan.Married()
	fmt.Println(Akhdan)
}