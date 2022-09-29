package main

import (
	"fmt"
)

func main() {

	var myName = "Brad"
	fmt.Println("My name is", myName)

	var aName string = "Tyler"
	fmt.Println("Name:", aName)

	userName := "Lukas"
	fmt.Println("Userame:", userName)

	var aSum int
	fmt.Println("The sume is:", aSum)

	part1, other := 1, 5
	fmt.Println("Part 1 is:", part1, "|", "Other is:", other)

	part2, other := 2, 0
	fmt.Println("Part 2 is:", part2, "|", "Other is:", other)

	aSum = part1 + part2
	fmt.Println("Sum is:", aSum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Println("Lesson Name is:", lessonName, "\nLesson Type is:", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)

}
