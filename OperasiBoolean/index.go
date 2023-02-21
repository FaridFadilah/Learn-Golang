package main

import "fmt"

func main(){
	var (
		nilaiAkhir = 90
		absensi = 80

		lulusNilaiAkhir bool = nilaiAkhir >= 80
		lulusAbsensi bool = absensi >= 80

		lulus bool = lulusNilaiAkhir && lulusAbsensi
	)

	fmt.Println(lulusNilaiAkhir)
	fmt.Println(lulusAbsensi)
	fmt.Println(lulus)
}