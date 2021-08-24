package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

func triangle() {
	const a, b = 3, 4
	fmt.Println(calcTriangle(a, b))

}

func calcTriangle(a, b int) int {
	c := int(math.Sqrt(float64(a*a + b*b)))
	return c

}

func enums() {
	const (
		a = iota * 10
		b
		c
	)
	fmt.Println(a, b, c)
}

func main() {
	euler()
	triangle()
	enums()
}
