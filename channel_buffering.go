// package main

// import (
// 	"fmt"
// )

// func main() {

// 	messages := make(chan string, 2)
// 	// go func() {

// 	// 	messages <- "test ping"
// 	// }()

// 	// messages <- "buffered"

// 	messages <- "channel"

// 	fmt.Println(<-messages)

// 	fmt.Println(<-messages)

// }
