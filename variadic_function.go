package main

import "fmt"


func sumAll(numbers ...int) int { // number sama dengan slice
	total := 0

	for _, number := range numbers{
		total += number
	}
	return total
}

func main(){
	fmt.Println(sumAll(10, 20, 10, 4, 5, 6))
	
	//jika data di dalam slice 
	numbers := []int {1, 3, 10, 30, 20}
	fmt.Println(sumAll(numbers...)) //pemanggilan variabel slice dengan titik tiga ...
	
}