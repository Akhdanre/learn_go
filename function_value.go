package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main(){
	goodbye := getGoodBye  //membuat variabel mejadi atau menampung fungsi
	fmt.Println(goodbye("Akhdan"))
}