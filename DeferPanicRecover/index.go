package main

import "fmt"

func logging(){
	fmt.Println("Selesai memanggil Logging")
}

func runLogging(value int){
	defer logging()
	fmt.Println("Jalankan aplikasi")
	result := 10 / value
	fmt.Println(result)
}

func endApp(){
	fmt.Println("Aplikasi selesai")
	message := recover()
	if message != nil{
		fmt.Println(message)
	}
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runLogging(5)
	runApp(false)
}