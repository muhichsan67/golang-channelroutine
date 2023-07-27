package go_routine

import (
	"fmt"
	"testing"
	"time"
)

// GO ROUTINE TIDAK BISA MERETURN VALUE
func RunHelloWorld(number int) {
	fmt.Println("Hello World", number)
}

func TestCreateGoroutine(t *testing.T) {
	for i := 0; i < 100; i++ {
		go RunHelloWorld(i + 1)
	}
	fmt.Println("Ups!")

	time.Sleep(1 * time.Second)
}
