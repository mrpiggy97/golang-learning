package catching

import (
	"fmt"
	"sync"

	"github.com/mrpiggy97/golang-learning/workerpools"
)

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	Value interface{}
	Err   error
}

type Memory struct {
	F      Function
	Cache  map[int]FunctionResult
	Locker *sync.Mutex
}

func NewCache(f Function) *Memory {
	return &Memory{
		F:      f,
		Cache:  make(map[int]FunctionResult),
		Locker: new(sync.Mutex),
	}
}

func (memInstance *Memory) Get(key int) (interface{}, error) {
	result, exists := memInstance.Cache[key]
	if !exists {
		fmt.Println("doing calculation")
		result.Value, result.Err = memInstance.F(key)
		memInstance.Cache[key] = result
	}
	return result.Value, result.Err
}

func GetFibonacci(n int) (interface{}, error) {
	return workerpools.Fibonacci(n), nil
}
