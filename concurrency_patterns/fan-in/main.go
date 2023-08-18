package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// fanIn combines multiple input channels into a single output channel.
// The Fan-In Pattern is a consolidation of multiple channels into one
// channel by multiplexing each received value.
func fanIn(cs ...<-chan []string) <-chan []string {
	chans := len(cs)
	wait := make(chan struct{}, chans)

	out := make(chan []string)

	send := func(c <-chan []string) {
		defer func() { wait <- struct{}{} }()

		for n := range c {
			out <- n
		}
	}

	for _, c := range cs {
		go send(c)
	}

	go func() {
		for range wait {
			chans--
			if chans == 0 {
				break
			}
		}

		close(out)
	}()

	return out
}

func readCSV(file string) (<-chan []string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("opening file %v", err)
	}

	ch := make(chan []string)

	cr := csv.NewReader(f)

	go func() {
		for {
			record, err := cr.Read()
			if err == io.EOF {
				close(ch)

				return
			}

			ch <- record
		}
	}()

	return ch, nil
}

// What is the Fan-In concurrency pattern?
// Consolidation of multiple channels into one channel by multiplexing each recieved value.
func main() {

	ch1, err := readCSV("file1.csv")
	if err != nil {
		panic(fmt.Errorf("could not read file1 %v", err))
	}

	ch2, err := readCSV("file2.csv")
	if err != nil {
		panic(fmt.Errorf("could not read file2 %v", err))
	}

	//-

	exit := make(chan struct{})

	ch := fanIn(ch1, ch2)

	go func() {
		for v := range ch {
			fmt.Println(v)
		}

		close(exit)
	}()

	<-exit

	fmt.Println("All completed, exiting")
}
