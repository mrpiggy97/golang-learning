package atomics

import (
	"sync/atomic"
	"time"
)

func Increaser(counter *int32) {
	for {
		atomic.AddInt32(counter, 2)
		time.Sleep(time.Millisecond * 500)
	}
}

func Decreaser(counter *int32) {
	for {
		atomic.AddInt32(counter, -1)
		time.Sleep(time.Millisecond * 250)
	}
}
