package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(EasyLine(7))
	fmt.Println(EasyLine(100))
}

func EasyLine(n int) string {
	var result = big.NewInt(0)

	for i := 0; i <= n; i++ {
		binomialCoefficient := big.NewInt(1).Binomial(int64(n), int64(i))
		result = result.Add(result, binomialCoefficient.Mul(binomialCoefficient, binomialCoefficient))
	}

	return result.String()
}
