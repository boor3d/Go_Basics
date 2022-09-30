//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func nameGreeting(name string) string {
	fmt.Println("Hello", name, "!")
	return ""
}

func wtf() string {
	return "I guess anything goes here"
}

func addThree(x, y, z int) int {
	return x + y + z
}

func aNum(n int) int {
	return n
}

func twoNum(x int, y int) (int, int) {
	return x, y
}

func main() {
	nameGreeting("Brad")
	fmt.Println(wtf())

	a, b := twoNum(6, 9)

	fmt.Println(addThree(addThree(1, 2, 3), aNum(a), aNum(b)))

}
