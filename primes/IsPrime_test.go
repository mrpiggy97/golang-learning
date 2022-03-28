package primes_test

import (
	"testing"

	"github.com/mrpiggy97/golang-learning/primes"
)

func checkingIfPrime(testCase *testing.T) {
	var number int64 = 13
	var isPrime bool = primes.IsPrime(number)
	if !isPrime {
		testCase.Errorf("%v is a prime number", number)
	}
}

func TestIsPrime(testCase *testing.T) {
	testCase.Run("action=test-is-prime", checkingIfPrime)
}
