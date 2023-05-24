package main

import "fmt"

func main(){
	type NIM string

	var ktmNF NIM = "666"
	fmt.Println(ktmNF)
	fmt.Println(NIM("999"))
}