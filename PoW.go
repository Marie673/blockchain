package main

import (
	"bytes"
	"math"
)

func compare(a [32]byte, b []byte) int {
	check := make([]byte, 32)
	for i := 0; i < 32; i++ {
		check[i] = a[i]
	}
	return bytes.Compare(b, check)
}

func makeByte(n int) []byte {
	ans := make([]byte, 32)
	if n >= 256 {
		ans[0] = byte(255)
		return ans
	}
	remind := n % 8
	quot := n / 8
	check := int(math.Pow(2, float64(remind)))
	ans[31-quot] = byte(check)

	return ans
}
