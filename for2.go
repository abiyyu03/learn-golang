package main

import "fmt"

//for kayak while
func main(){
	number := 1
	for number < 10 {
		fmt.Println("Loop ke", number)
		if number == 5 {
			break
		}
		number++
	}
}