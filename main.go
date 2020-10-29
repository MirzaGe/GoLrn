package main

import "fmt"

func main() {
	x := 11 % 3
	fmt.Println(x)
	if x == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

	for i := 1; i < 9; i++ {
		if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}

}

// constants,
