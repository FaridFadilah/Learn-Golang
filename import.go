package main

import (
	accessmodifier "Golang-Dasar/AccessModifier"
	"Golang-Dasar/PackageImport"
	packageinit "Golang-Dasar/PackageInit"
	"fmt"
)

func main(){
	PackageImport.SayHello("farid")

	accessmodifier.SayGoodbye()
	// accessmodifier.sayGood() // error

	fmt.Println(accessmodifier.Apps)
	// fmt.Println(accessmodifier.ver) //error

	packageinit.Database()
}