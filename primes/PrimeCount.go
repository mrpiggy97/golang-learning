package primes

func PrimeCount(number int) int {
	var count int = 0
	for i := 0; i <= number; i++ {
		var isPrime bool = IsPrime(int64(i))
		if isPrime {
			count = count + 1
		}
	}
	return count
}
