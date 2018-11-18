package main

import "fmt"

// func main() {
// 	fmt.Print("Hello world")
// }
func main() {
	fmt.Println(Add(21, 21))
}
func Add(a, b int) int {
	return a + b
}
