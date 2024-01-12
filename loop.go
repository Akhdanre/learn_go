package main


import "fmt"

func main(){
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("loop in", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("loop in", counter)
	}

	names := []string{"Akhdan", "Robbani"}
	for i := 0; i < len(names); i++{
		fmt.Println(names[i])
	}	

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
	for _, name := range names {
		fmt.Println( "value", name)
	}
}