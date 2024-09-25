package prac1

import (
	"fmt"
	"math"
)

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Run() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Please enter a non-negative number.")
	} else {
		fact := factorial(num)
		prime := isPrime(num)

		switch {
		case prime:
			fmt.Printf("%d is a prime. Its factorial is %d.\n", num, fact)
		case num < 10:
			fmt.Printf("%d is not prime but less than 10. Its factorial is %d.\n", num, fact)
		default:
			fmt.Printf("%d is not prime. Its factorial is %d.\n", num, fact)
		}
	}
}
