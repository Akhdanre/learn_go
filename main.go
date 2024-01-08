package main

import "fmt"
import "rsc.io/quote"

func main(){
	fmt.Println("Hello, World !")
	fmt.Println(quote.Go())
	fmt.Println(Hello("akhdan"))
}


func Hello(name string) string{
	message := fmt.Sprintf("HI %v, Welcome", name)
	return message
}