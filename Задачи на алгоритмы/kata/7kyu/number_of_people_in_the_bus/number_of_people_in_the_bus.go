package main

func Number(busStops [][2]int) int {
	in, out := 0, 0

	for _, v := range busStops {
		in += v[0]
		out += v[1]
	}

	return in - out
}
