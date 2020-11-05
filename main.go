package main

import "fmt"

func one(x *int) {
	*x = 1
}
func main() {
	y := new(int)
	one(y)
	fmt.Println(*y)
}
func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}
