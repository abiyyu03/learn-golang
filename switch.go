package main

import "fmt"

func main(){
	name := "Abi"

	switch name {
	case "Abi":
		fmt.Println("Hello Abi")
	case "Fulan":
		fmt.Println("Hello Fulan")
	default:
		fmt.Println("What is ur name ?")
	}
}