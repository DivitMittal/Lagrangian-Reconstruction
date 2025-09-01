package main

import "math/big"

func lagrangeInterpolation(points []Point) *big.Int {
	resultNum := big.NewInt(0)
	resultDen := big.NewInt(1)
	n := len(points)

	for i := 0; i < n; i++ {
		numerator := big.NewInt(1)
		denominator := big.NewInt(1)

		for j := 0; j < n; j++ {
			if i != j {
				numerator.Mul(numerator, big.NewInt(0).Neg(points[j].X))
				diff := big.NewInt(0).Sub(points[i].X, points[j].X)
				denominator.Mul(denominator, diff)
			}
		}

		termNum := big.NewInt(0).Mul(points[i].Y, numerator)
		termNum.Mul(termNum, resultDen)

		resultNumPart := big.NewInt(0).Mul(resultNum, denominator)
		resultNum.Add(resultNumPart, termNum)
		resultDen.Mul(resultDen, denominator)

		gcd := big.NewInt(0).GCD(nil, nil, resultNum, resultDen)
		resultNum.Div(resultNum, gcd)
		resultDen.Div(resultDen, gcd)
	}

	if resultDen.Cmp(big.NewInt(-1)) == 0 {
		return big.NewInt(0).Neg(resultNum)
	}

	if resultDen.Cmp(big.NewInt(1)) != 0 {
		return big.NewInt(0).Div(resultNum, resultDen)
	}

	return resultNum
}
