package main

import (
	"flag"
	"fmt"
	"math/big"
)

func main() {
	aStr := flag.String("a", "0", "set `a` value")
	bStr := flag.String("b", "0", "set `b` vblue")
	flag.Parse()

	a := big.NewInt(0)
	if err := a.UnmarshalText([]byte(*aStr)); err != nil {
		panic(err)
	}
	b := big.NewInt(0)
	if err := b.UnmarshalText([]byte(*bStr)); err != nil {
		panic(err)
	}

	fmt.Printf("Multiplication: %s\n", big.NewInt(0).Mul(a, b))
	fmt.Printf("Integer division: %s\n", big.NewInt(0).Div(a, b))
	aFloat, _, _ := big.ParseFloat(*aStr, 10, big.MaxPrec, big.ToNearestEven)
	bFloat, _, _ := big.ParseFloat(*bStr, 10, big.MaxPrec, big.ToNearestEven)
	fmt.Printf("Floating division: %s\n", big.NewFloat(0).Quo(aFloat, bFloat).String())
	fmt.Printf("Addition: %s\n", big.NewInt(0).Add(a, b))
	fmt.Printf("Subtraction: %s\n", big.NewInt(0).Sub(a, b))
}
