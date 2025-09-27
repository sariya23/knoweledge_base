package pop

func PopCountViaSum(pc [256]byte, x uint64) int {
	return int(
		pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func PopCountViaFor(pc [256]byte, x uint64) int {
	var sum int

	for i := 1; i < 8; i++ {
		sum += int(pc[byte(x>>i*8)])
	}
	return sum
}
