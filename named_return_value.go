package main


import "fmt"

func getCompleteName() (firstname, lastname string){
	firstname = "akhdan"
	// lastname = "robbani"
	return firstname, lastname
}

func main(){
	firstname, lastname := getCompleteName()
	fmt.Println(firstname, lastname)
}