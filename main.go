package main

import "fmt"

func main() {
	foo(2, 3, 4, 5, 5, 6)
}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding,", v, sum)
	}

	fmt.Println("The total is,", sum)
}
