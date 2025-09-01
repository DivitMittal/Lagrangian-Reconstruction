package main

import (
	"fmt"
	"math/big"
	"strings"
)

func convertFromBase(value string, base int) (*big.Int, error) {
	result := big.NewInt(0)
	baseInt := big.NewInt(int64(base))

	for _, char := range strings.ToUpper(value) {
		var digit int64
		if char >= '0' && char <= '9' {
			digit = int64(char - '0')
		} else if char >= 'A' && char <= 'Z' {
			digit = int64(char - 'A' + 10)
		} else {
			return nil, fmt.Errorf("invalid character '%c'", char, base)
		}

		if digit >= int64(base) {
			return nil, fmt.Errorf("digit %d isn't possible for base %d", digit, base)
		}

		result.Mul(result, baseInt)
		result.Add(result, big.NewInt(digit))
	}

	return result, nil
}
