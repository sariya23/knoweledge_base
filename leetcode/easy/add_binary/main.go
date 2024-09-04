package main

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) string {
	a, b = addBitsForSameLength(a, b)
	p := false
	res := ""

	for i := len(a) - 1; i > -1; i-- {
		bit1 := binBoolMapper[string(a[i])]
		bit2 := binBoolMapper[string(b[i])]
		moduleSum, p1 := fullAdder(bit1, bit2, p)
		fmt.Printf("sum for bit1=%s, bit2=%s is %t. P is %t\n", string(a[i]), string(b[i]), moduleSum, p1)
		res = boolBinMapper[moduleSum] + res
		p = p1
	}

	if p {
		res = "1" + res
	}
	return res
}

func fullAdder(a, b, cIn bool) (bool, bool) {
	moduleSum := xor(xor(a, b), cIn)
	cOut := (a && b) || (cIn && (xor(a, b)))
	return moduleSum, cOut
}

func xor(a, b bool) bool {
	return (a && !b) || (!a && b)
}

func addBitsForSameLength(a, b string) (string, string) {
	if len(a) > len(b) {
		bitsToAdd := strings.Repeat("0", len(a)-len(b))
		return a, bitsToAdd + b
	} else if len(a) < len(b) {
		bitsToAdd := strings.Repeat("0", len(b)-len(a))
		return bitsToAdd + a, b
	}
	return a, b
}

var binBoolMapper = map[string]bool{
	"1": true,
	"0": false,
}

var boolBinMapper = map[bool]string{
	true:  "1",
	false: "0",
}

func main() {

}
