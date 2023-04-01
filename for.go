// package main

// import "fmt"

// func main() {
// 	fmt.Println("Tìm ước")
// 	var n = 100

// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			fmt.Println("-> ", i)
// 		}
// 		if n%(n/i) == 0 && n/i != i {
// 			fmt.Println("-> ", n/i)
// 		}
// 	}

// }

// package main

// import "fmt"

// func main() {
// 	i := 1

// 	fmt.Println("case 1: ")

// 	for i <= 3 {
// 		fmt.Println(i)
// 		i++
// 	}

// 	fmt.Println("case 2: ")
// 	for i := 1; i < 6; i++ {
// 		fmt.Println(i)
// 	}

// 	for {
// 		fmt.Println("break while")
// 		break
// 	}

// 	for i := 0; i <= 6; i++ {
// 		if i&1 == 1 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// }
