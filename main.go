package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 8, 55, 66, 77, 777, 7545}
	xs := []string{"james", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("----------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
