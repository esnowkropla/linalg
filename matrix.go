package main

import (
	"fmt"
)

type Matrix struct {
	row, col uint
	Mij []float64
}

func main() {
	A := new(Matrix)
	fmt.Printf("%v\n", A)
}
