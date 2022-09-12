package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewFloat(6e21)
	b := big.NewFloat(3e21)
	x := new(big.Float)
	x.Mul(a, b)
	fmt.Printf("%f * %f = %f\n", a, b, x)
	x.Add(a, b)
	fmt.Printf("%f + %f = %f\n", a, b, x)
	x.Sub(a, b)
	fmt.Printf("%f - %f = %f\n", a, b, x)

	c := new(big.Int)
	c.SetString("600000000000000000000000", 10)
	d := new(big.Int)
	d.SetString("300000000000000000000000", 10)
	z := new(big.Int)
	z.Div(c, d)
	fmt.Printf("%d / %d = %d\n", c, d, z)
}
