package main

import "fmt"

func main() {
	var a, b, c, d int
	var e float64
	a = 10
	b = 2
	c = b * a
	d = c + a
	e = d / a

	fmt.Printf("value of c %d\n", c)
	fmt.Printf("value of d %d\n", d)
	fmt.Printf("value of e %f\n", e)

}
