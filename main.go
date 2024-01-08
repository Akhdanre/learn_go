package main

import "fmt"
import "rsc.io/quote"

func main(){
	fmt.Println("Hello, World !")
	fullname := "Akhdan Robbani"
	fmt.Println(hello(fullname))
	fmt.Println(quote.Go())
}


func hello(name string) string{
	message := fmt.Sprintf("%v in here", name)
	return message
}