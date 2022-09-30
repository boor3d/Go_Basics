package main

import "fmt"

func double(x int) int {
	return x * 2
}

func add(x int, y int) int {
	return x + y
}

func greet() {
	fmt.Println("Hello From Greet Function!")
}

func main() {
	greet()

	dub := double(4)

	adder := add(dub, 5)

	fmt.Println("Doubled value is:", dub, "\nAdded Function is:", adder)

	anotherNumber := add(double(5), 100)

	fmt.Println("This Random Number is:", anotherNumber)

}
