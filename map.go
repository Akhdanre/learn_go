package main


import "fmt"


func main(){
	//cara pertama
	var user map[string]string = map[string]string{}
	user["name"] = "Akhdan Robbani"
	user["address"] = "Nganjuk"

	// cara lain
	user2 := map[string]string{
		"name" : "oukenzuemasio",
		"address" : "Nganjuk",
	}
	fmt.Println(user2)


	book := make(map[string]string)
	book["title"] = "learn_golang"
	book["author"] = "Akhdan"
	book["ups"] = "typo"

	fmt.Println(book)
	
	delete(book, "ups")
	fmt.Println(book)
}