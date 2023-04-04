package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {

	fmt.Print("working...")

	time.Sleep(time.Second)

	fmt.Println("done")

	done <- true
	fmt.Println("finished...............")
}

func main() {

	done := make(chan bool)
	go worker(done)
	time.Sleep(time.Second)
	go worker(done)

	<-done
	<-done
}
