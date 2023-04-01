// package main

// import (
// 	"fmt"
// )

// func Fake_Map[K, V any](s []K, transform func(K) V) []V {
// 	rs := make([]V, len(s))
// 	for i, v := range s {
// 		rs[i] = transform(v)
// 	}
// 	return rs
// }

// func main() {
// 	arr := []int{1, 2, 3}
// 	resultArr := Fake_Map(arr, func(v int) int { return v * 2 })
// 	fmt.Println(resultArr)
// }


