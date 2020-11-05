package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 6
}
func main() {
	x := 7
	zero(&x)
	fmt.Println(x)
}

//
