package main

import "fmt"

func main(){
	name := "Abiyyu bin Fulan"
	length := len(name)

	switch {
	case length > 20:
		fmt.Println("Name too long")
	case length > 10:
		fmt.Println("quite a long name")
	default:
		fmt.Println("Name is correct")
	}
}