package main
import "fmt"

func main(){
	names := [3]string{"Abiyyu Cakra","Ayyas Mumtaz","Daffa Fawwaz"}
	myslice := names[0:1]

	fmt.Println("My name is ", myslice)
}