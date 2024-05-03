package main

import (
    "fmt"
    "math"
)

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    if n <= 3 {
        return true
    }
    if n%2 == 0 || n%3 == 0 {
        return false
    }
    i := 5
    for i*i <= n {
        if n%i == 0 || n%(i+2) == 0 {
            return false
        }
        i += 6
    }
    return true
}

func primeNumbersInRange(min, max int) []int {
    primes := []int{}
    for i := min; i <= max; i++ {
        if isPrime(i) {
            primes = append(primes, i)
        }
    }
    return primes
}

func main() {
    min := 11
    max := 20
    primes := primeNumbersInRange(min, max)
    fmt.Println(primes)
}
