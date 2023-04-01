// package main

// import (
// 	"fmt"
// )

// func main() {

// 	fmt.Println("Running........")

// 	go_map := make(map[string]int)

// 	go_map["1"] = 1

// 	go_map["2"] = 2

// 	fmt.Println(go_map)

// 	fmt.Println(go_map["2"])

// 	go_map["add"] = 3

// 	fmt.Println(go_map)

// 	delete(go_map, "add")

// 	fmt.Println(go_map)

// 	err, prs := go_map["k2"]
// 	fmt.Println("_:", err)
// 	fmt.Println("prs:", prs)

// 	new_map := map[string]int{}
// 	fmt.Printf("new_map: %v", new_map)
// 	fmt.Printf("\n%T\n", new_map)
// }
