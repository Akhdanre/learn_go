package main

import "fmt"

func logging(){
	fmt.Println("selesai memanggil function")
}


func runApplication(){
	// defer akan dieksekusi di akhir kode
	defer logging() //lebih baik menggunakan defer daripada mengatur posisi di bawah
	fmt.Println("akhir kode")
}

func main(){
	runApplication()
}