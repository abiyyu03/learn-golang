package main

import "fmt"

func main(){
	person := map[string]string{"name":"Abiyyu", "address":"Jawa Barat"}

	fmt.Println(person)
	fmt.Println("My name is ", person["name"])
	fmt.Println("My address is ", person["address"])
}