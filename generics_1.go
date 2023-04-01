package main

import "fmt"

func Printtt[T any](s []T) {
	for index, value := range s {
		fmt.Print(index)
		fmt.Print("-->")
		fmt.Print(value)
		fmt.Print("\n")

	}
}
func main() {
	Printtt([]string{"Hello, ", "playground\n"})
	Printtt([]int{1, 2, 3})
}
