// package main

// import (
// 	"fmt"
// )

// type person struct {
// 	name string
// 	age  int
// }

// func newPerson(name string) *person {

// 	p := person{name: name}
// 	p.age = 42
// 	return &p
// }

// func main() {

// 	fmt.Println("Running........")

// 	fmt.Println(person{"Bob", 20})

// 	fmt.Println(person{name: "hanb", age: 23})

// 	fmt.Println(person{name: "hanb"})

// 	fmt.Println(&person{name: "hanb", age: 23})

// 	fmt.Println(".....................")

// 	fmt.Println(newPerson("hanb"))

// 	s := person{name: "hanbbb", age: 50}

// 	fmt.Println("name: ", s.name)

// 	sp := &s

// 	fmt.Println("age: ", sp.age)

// 	sp.age = 90

// 	fmt.Println("new age: ", sp.age)

// 	fmt.Println("old: ", s)
// }
