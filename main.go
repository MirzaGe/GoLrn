package main

import "fmt"

func main() {
	foo()
	bar("james")
	S1 := woo("MoneyPenny")
	fmt.Println(S1)
	fmt.Println("Hello" + "!" + "!")
}
func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from woo,", st)
}
