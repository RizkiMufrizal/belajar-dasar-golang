package belajargolangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestCreateContext(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)

	fmt.Println(contextB.Value("b"))
	fmt.Println(contextD.Value("d"))
	fmt.Println(contextD.Value("b"))
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destionation := CreateCounter(ctx)
	for n := range destionation {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}
	cancel()

	time.Sleep(2 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destionation := CreateCounter(ctx)
	for n := range destionation {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println(runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destionation := CreateCounter(ctx)
	for n := range destionation {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println(runtime.NumGoroutine())
}
