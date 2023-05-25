package main

import "fmt"

// for range : looping untuk array / map
func main(){
	names := []string{"Abiyyu","Fulan"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
	person := map[string]string{"Abi":"Jawa Barat", "Rio":"Jawa Tengah","dede":"Banten"}
	for name, address := range person {
		fmt.Println("name :", name, " address :", address)
	}
}