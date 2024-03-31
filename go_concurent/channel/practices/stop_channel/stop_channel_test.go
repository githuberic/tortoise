package stop_channel

import (
	"fmt"
	"testing"
	"time"
)

func worker(stopCh <-chan struct{}) {
	go func() {
		defer func() {
			fmt.Println("Worker exit")
		}()

		t := time.NewTicker(time.Millisecond * 500)

		// Using stop channel explicit exit
		for {
			select {
			case <-stopCh:
				fmt.Println("Recv stop signal")
				return
			case <-t.C:
				fmt.Println("Working .")
			}
		}
	}()
	return
}

func TestCloseChannel(t *testing.T)  {
	stopCh := make(chan struct{})
	worker(stopCh)

	time.Sleep(time.Second * 2)
	close(stopCh)


	// Wait some print
	time.Sleep(time.Second)
	fmt.Println("main exit")
}