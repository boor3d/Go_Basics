package main

import "fmt"

func price(x int) int {
	return x
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {

	switch p := price(7); {
	case p < 2:
		fmt.Println("Cheap Item")

	case p < 10:
		fmt.Println("Moderatley Priced Item")

	default:
		fmt.Println("Expensive Item")
	}

	ticket := 7
	switch ticket {
	case Economy:
		fmt.Println("Economy Seating")

	case Business:
		fmt.Println("Business Seating")

	case FirstClass:
		fmt.Println("Flyin' First Class, up in the sky!")

	default:
		fmt.Println("LOL Can't even get on the plane!")
	}

}
