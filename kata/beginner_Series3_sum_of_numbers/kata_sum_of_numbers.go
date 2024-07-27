package kata

func GetSum(a, b int) int {
	start := a
	stop := b

	if b < a {
		start = b
		stop = a
	}

	total := 0

	for i := start; i < stop+1; i++ {
		total += i
	}
	return total
}
