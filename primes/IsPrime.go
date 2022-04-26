package primes

func IsPrime(number int64) bool {
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
