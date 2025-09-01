package main

import "math/big"

type Point struct {
	X, Y *big.Int
}

type TestCase struct {
	Keys   Keys                         `json:"keys"`
	Shares map[string]map[string]string `json:"-"`
}

type Keys struct {
	N int `json:"n"`
	K int `json:"k"`
}
