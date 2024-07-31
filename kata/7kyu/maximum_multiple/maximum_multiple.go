package main

func MaxMultiple(d, b int) int {
	dividend := 0
	for i := d; i < b+1; i++ {
		if i%d == 0 {
			dividend = i
		}
	}
	return dividend
}
