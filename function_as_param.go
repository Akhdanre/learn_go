package main

import "fmt"

//menyederhanakan function dengan alias
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter){
	filteredName := filter(name)
	fmt.Println("hai " + filteredName)
}

func filterName(name string) string{
	if name == "anjing" {
		return ""
	}else {
		return name
	}
}

func main(){
	sayHelloWithFilter("akeon", filterName)
}