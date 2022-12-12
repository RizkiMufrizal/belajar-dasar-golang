package belajargolanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "rizki mufrizal"
		fmt.Println("selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func giveMeResponse(channel chan string) {
	channel <- "krim ke channel"
}

func TestCreateChannelParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go giveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
}

func TestCreateChannelBuffer(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	channel <- "rizki"
	channel <- "mufrizal"
	fmt.Println(len(channel))

	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println(len(channel))
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}
	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go giveMeResponse(channel1)
	go giveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			counter++
			fmt.Println("data dari channel 1", data)
		case data := <-channel2:
			counter++
			fmt.Println("data dari channel 2", data)
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}
