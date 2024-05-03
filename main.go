package main

import (
	"fmt"
)

func computerInflection(num int) string {
	if num == 1 {
		return fmt.Sprintf("%d computer", num)
	}
	return fmt.Sprintf("%d computers", num)
}

func main() {
	var input int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println(computerInflection(input))
}
