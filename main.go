package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 0.3048
	fmt.Println(output)
}

/* Using the example program as a starting point,
write a program that converts
from Fahrenheit into Celsius (C = (F − 32) * 5/9). */
