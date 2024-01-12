package main

import "fmt"


type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name){
		fmt.Println("you're name has blocked", name)
	}else {
		fmt.Println("Welcome", name)
	}
}

func main(){
	//anonym di simpan di variabel
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("anjing", blacklist)
	registerUser("akeon", func(name string) bool { // anonym langsung di dalam parameter
		return name == "anjing"
	})
}