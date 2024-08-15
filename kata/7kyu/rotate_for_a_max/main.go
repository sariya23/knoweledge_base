package main

import (
	"fmt"
	"strconv"
)

func Rotate(n int, offset int) int {
	stringRep := strconv.Itoa(n)

	digitToRemove := stringRep[offset]
	rotated := stringRep[:offset] + stringRep[offset+1:] + string(digitToRemove)
	rotatedNumber, _ := strconv.Atoi(rotated)

	return rotatedNumber
}

func MaxRot(n int64) int64 {
	maxRotate := n

	for offset := 0; offset < len(strconv.Itoa(int(n)))-1; offset++ {
		n = int64(Rotate(int(n), offset))
		if n > int64(maxRotate) {
			maxRotate = n
		}
	}

	return int64(maxRotate)
}

func main() {
	// n := 56789
	// for i := 0; i < len(strconv.Itoa(n)); i++ {
	// 	n = Rotate(n, i)
	// 	fmt.Println(n)
	// }
	fmt.Println(MaxRot(56789))
}
