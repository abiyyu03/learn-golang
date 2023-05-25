package main

import "fmt"

func main() {
	var number int = 5
	var numberPointer *int = &number

	fmt.Println("NIlai dari number:", number)
	fmt.Println("Alamat Memori dari number:", &number)
	fmt.Println("Nilai dari number pointer:", *numberPointer)
	fmt.Println("Alamat memori dari number pointer:", numberPointer)
}