package main

import "fmt"

func makeEvenGenerator() func() uint{
	i := uint(0)
	return func() (return uint){
		ret = i
		i += 2
		return makeEvenGenerator()

	}
}
