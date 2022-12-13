package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}
	group := sync.WaitGroup{}

	pool.Put("rizki")
	pool.Put("mufrizal")

	for i := 0; i < 10; i++ {
		go func() {
			defer group.Done()
			data := pool.Get()
			fmt.Println("data", data)
			pool.Put(data)

			group.Add(1)
		}()
	}

	group.Wait()
	fmt.Println("Selesai")
}

var data sync.Map
var addToMap = func(value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 1; i <= 100; i++ {

		go addToMap(i, &group)
	}

	group.Wait()
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 1; i <= 10; i++ {
		group.Add(1)
		go WaitCondition(i)
	}

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()

	group.Wait()
}
