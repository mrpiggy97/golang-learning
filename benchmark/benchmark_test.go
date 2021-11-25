package benchmark

import (
	"fmt"
	"testing"
)

func Sum(number int64) int64 {
	var i int64 = 0
	var result int64 = 0
	for i <= number {
		result = result + i
		i++
	}
	return result
}

func BenchmarkSum(benchMarkCase *testing.B) {
	//this is a basic benchmark test
	fmt.Println("benchmark n:", benchMarkCase.N)
	for i := 0; i < benchMarkCase.N; i++ {
		Sum(100000)
	}
}

func BenchmarkSumParallel(benchmarkCase *testing.B) {
	benchmarkCase.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Sum(100000)
		}
	})
}
