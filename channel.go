// package main

// import (
// 	"fmt"
// )

// func main() {

// 	messages := make(chan string)

// 	go func() {

// 		messages <- "test ping"
// 	}()

// 	receives := <-messages

// 	fmt.Println("receives: ", receives)

// }
