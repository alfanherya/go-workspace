package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Alfan herya"

		fmt.Println("Selesai mengirim data ke channel")
	}()

	data := <-channel
	fmt.Print(data)
	time.Sleep(5 * time.Second)

}
