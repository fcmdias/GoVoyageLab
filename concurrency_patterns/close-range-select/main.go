package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	exit := make(chan struct{})

	// Method 1 using "select"
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(time.Now(), i, "sending")
			ch <- i
			fmt.Println(time.Now(), i, "sent")
			time.Sleep(1 * time.Second)
		}

		fmt.Println(time.Now(), "all completed, leaving")
		close(ch)
	}()

	// Method 2 using "range"
	// go func() {
	// 	for {
	// 		select {
	// 		case v, open := <-ch:
	// 			if !open {
	// 				close(exit)
	// 				return
	// 			}
	// 			fmt.Println(time.Now(), "received", v)
	// 		default:
	// 			fmt.Println("nothing is happening")
	// 		}
	// 	}
	// }()

	go func() {
		for v := range ch {
			fmt.Println(time.Now(), "received", v)
		}
		close(exit)
	}()

	fmt.Println(time.Now(), "waiting for everything to complete")
	<-exit
	fmt.Println(time.Now(), "exiting")

}
