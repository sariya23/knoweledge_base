package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 {
			n--
		}
		return n <= 0
	}

	if flowerbed[0] == 0 && flowerbed[1] != 1 {
		flowerbed[0] = 1
		n--
	}
	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] != 1 {
			flowerbed[i] = 1
			n--
		}
	}

	if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
		n--
	}

	return n <= 0
}
