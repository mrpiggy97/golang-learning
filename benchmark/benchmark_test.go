package benchmark

import (
	"testing"
)

func isPrime(number int) bool {
	if number == 2 || number == 3 || number == 5 || number == 7 || number == 9 {
		return true
	}
	if number%2 == 0 {
		return false
	}
	if number%3 == 0 {
		return false
	}
	if number%5 == 0 {
		return false
	}
	if number%7 == 0 {
		return false
	}
	if number%9 == 0 {
		return false
	}
	return true
}

func primeCount(number int) int {
	var count int = 0
	for i := 2; i <= number; i++ {
		var numberIsPrime bool = isPrime(i)
		if numberIsPrime {
			count = count + 1
		}
	}
	return count
}

func BenchmarkPrimeCount(benchCase *testing.B) {
	for i := 0; i <= benchCase.N; i++ {
		primeCount(i)
	}
}
