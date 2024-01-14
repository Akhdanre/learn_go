package helper

import "fmt"

//diawali huruf kecil tidak akan bisa di akses di luar package
var version = "1.0.0"
//diawali huruf besar bisa diakses di luar package
var Application = "learn go"

//sama dengan variabel, fungsi jika diawali huruf besar bisa diakses di luar package
func SayHello(name string) string {
	return "Hello, " + name
}
//tidak bisa diakses di luar package 
func sayGodBye(name string) string {
	return "good bye "+ name
}

func Example(){
	//bisa di akses di fungsi dengan catatan di dalam package yang sama
	sayGodBye("akhdan")
	fmt.Println(version)
}

