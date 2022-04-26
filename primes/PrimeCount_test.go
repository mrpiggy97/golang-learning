package primes_test

import (
	"testing"

	"github.com/mrpiggy97/golang-learning/primes"
)

func TestPrimeCount(testCase *testing.T) {
	var limit int = 12323412
	var count int = primes.PrimeCount(limit)
	if count <= 0 {
		testCase.Error("PrimeCount does not count a single prime")
	}
}
