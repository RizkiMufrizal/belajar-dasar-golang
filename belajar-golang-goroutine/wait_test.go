package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 1; i < 100; i++ {
		go RunAsynchronous(&group)
	}

	group.Wait()
	fmt.Println("Selesai")
}

var counter = 0

func OnlyOnce() {
	counter++
}

func TestWaitOnce(t *testing.T) {
	var group sync.WaitGroup
	var once sync.Once

	for i := 1; i < 100; i++ {
		go func() {
			defer group.Done()

			group.Add(1)
			once.Do(OnlyOnce)
		}()
	}

	group.Wait()
	fmt.Println("counter", counter)
}
