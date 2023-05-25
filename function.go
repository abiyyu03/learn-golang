package main

import "fmt"

func sayHello(firstName string, lastName string){
	fmt.Printf("Hello, %s %s!\n", firstName, lastName)
}

func getHello(name string) string{
	return "Hello " + name
}

func main(){
	result := getHello("Fulan")
	sayHello("Abiyyu","Fulan")
	fmt.Println(result)
}