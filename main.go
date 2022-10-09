package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	value := 100000000
	// value := 100
	fmt.Printf("Find Primes under %d", value)
	fmt.Println()

	startTime := time.Now()

	primes := []int{}
	for i := 2; i < value; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	fmt.Println()
	fmt.Printf("The Prime number count is %d", len(primes))

	fmt.Println()
	elapsedTime := time.Since(startTime)
	fmt.Printf("실행시간: %s\n", elapsedTime)

}

func isPrime(number int) bool {
	// fmt.Printf("is number %d Prime?", number)

	if number < 2 {
		// fmt.Println()
		return false
	}
	if number == 2 {
		// fmt.Print(" It is Prime")
		// fmt.Println()
		return true
	}
	if number%2 == 0 {
		// fmt.Println()
		return false
	}

	rootOfNumber := int(math.Sqrt(float64(number)))
	// fmt.Printf(" root of number is %d", rootOfNumber)
	for i := 3; i <= rootOfNumber; i += 2 {
		if number%i == 0 {
			return false
		}
	}

	// fmt.Print(" It is Prime")
	// fmt.Println()
	return true
}
