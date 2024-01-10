package main

import "fmt"

func main(){
	names := [...]string {"akhdan", "robbani", "akeon", "oukenzeumasio", "ouken"}
	
	//get value from index 3 until last index with slice
	slice1 := names[3:]
	fmt.Println(slice1)


	//get value from index 0 until index 2, that means 3 value will in slice
	slice2 := names[0:2]
	fmt.Println(slice2)


	// get length value in slice
	fmt.Println(len(slice1))


	days := [...]string {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	dayslice1 := days[5:]
	//this method will make parent array update to with slice
	dayslice1[0] = "sabtu_baru"
	dayslice1[1] = "minggu_baru"
	
	fmt.Println("\nvalue days array :",days)
	fmt.Println("value daysslice1 :",dayslice1)

	fmt.Println("\nAppend")

	dayslice2 := append(dayslice1, "libur baru")
	dayslice2[0] = "sabtu_lama"

	fmt.Println(dayslice1)
	fmt.Println(dayslice2)
	fmt.Println(days)

	
}