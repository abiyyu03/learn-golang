package main

import "fmt"

func main(){
	name := "Fulan"
	if name == "Fulan"	 {
		fmt.Println("Hello", name)
	} else if name == "Abi" {
		fmt.Println("Hello Abi")
	} else {
		fmt.Println("Eh salah orang, maaf") 
	}
}