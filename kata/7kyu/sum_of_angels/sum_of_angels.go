package main

func Angle(n int) int {
	baseAnglesSum := 360
	baseSides := 3

	return (n-baseSides)*180 + baseAnglesSum
}
