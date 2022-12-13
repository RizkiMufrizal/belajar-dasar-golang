package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	group := sync.WaitGroup{}
	var counter int64 = 0

	for i := 1; i <= 100; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	group.Wait()
	fmt.Println("counter", counter)
}
