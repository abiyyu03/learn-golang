package main
import "fmt"

func main(){
	const firstName = "Fulan"
	const lastName = "Cakra"

	const (
		age = 25
		bornPlace = "Tokyo"
	)
	//will error if redreclare the constant
	// firstName = "Fulan"
	// lastName = "fulan"

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
	fmt.Println(bornPlace)
}