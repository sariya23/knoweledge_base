package main

func ClosestMultipleOf10(n uint32) uint32 {
	lastDigit := n % 10
	if lastDigit >= 5 {
		return n + 10 - lastDigit
	}
	return n - lastDigit
}
