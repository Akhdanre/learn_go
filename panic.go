package main

import "fmt"

func endApp(){
	fmt.Println("end app")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(Error bool){
	defer endApp() //defer akan tetapi dijalankan meskipun terdapat error di panic
	if Error {
		panic("ups error")
	}
	//endapp() lebih baik menggunakan defer daripada opsi pemanggilan di akhir kode
}

func main(){
	runApp(true)
	fmt.Println("lanjut ke progmram selanjurnya")
}