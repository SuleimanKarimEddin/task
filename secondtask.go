package main

import (
    "fmt"
)

func findCommonDivisors(numbers []int) []int {
    if len(numbers) == 0 {
        return nil
    }

    // Find the minimum number in the array
    minNum := numbers[0]
    for _, num := range numbers {
        if num < minNum {
            minNum = num
        }
    }

    // Find common divisors
    var commonDivisors []int
    for i := 1; i <= minNum; i++ {
        isCommonDivisor := true
        for _, num := range numbers {
            if num%i != 0 {
                isCommonDivisor = false
                break
            }
        }
        if isCommonDivisor {
            commonDivisors = append(commonDivisors, i)
        }
    }

    return commonDivisors
}

func main() {
    numbers := []int{42, 12, 18}
    commonDivisors := findCommonDivisors(numbers)
    fmt.Println("Common divisors:", commonDivisors)
}
