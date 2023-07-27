package go_routine

import (
	"fmt"
	"testing"
	"time"
)

func GiveMeResponse(channel chan string) {

}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Ichsan Fathurrochman"
}

func OnlyOut(channel <-chan string) {
	nama := <-channel
	fmt.Println(nama)
}

func TestBufferedChannel(t *testing.T) {
	// Jumlah data yg bisa dimasukkan kedalam channel
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Muhammad"
	channel <- "Ichsan"
	channel <- "Ichsan"

	fmt.Println(cap(channel)) // melihat max kapasitas buffered channel
	fmt.Println(len(channel)) // melihat jumlah data buffered channel

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Selesai")
	// time.Sleep(5 * time.Second)
}
