package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	timeNew := <-timer.C
	fmt.Println(timeNew)

	channel := time.After(5 * time.Second)
	fmt.Println(<-channel)
}

func TestAfterTime(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("success run after fun")
		group.Done()
	})

	group.Wait()
}
