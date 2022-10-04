package main

import "fmt"

func average(a, b int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b) / 2
}

func main() {
	quiz1, quiz2 := 9, 8

	if quiz1 > quiz2 {
		fmt.Println("Quiz 1 is greater than quiz 2")
	} else if quiz1 < quiz2 {
		fmt.Println("Quiz 1 is less than quiz 2")
	} else {
		fmt.Println("The quiz scores are equal")
	}

}
