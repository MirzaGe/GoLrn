package main

import "fmt"

func zero(xyz *int) {
	*xyz = 6
}
func main() {
	x := 7
	zero(&x)
	fmt.Println(x)
}

//
