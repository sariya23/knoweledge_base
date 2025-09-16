package main

import (
	"fmt"
	"strconv"
)

func makeLettersKeyMap() map[string]int {
	m := map[string]int{}
	counter := 1
	for i := 97; i < 123; i++ {
		m[string(rune(i))] = counter
		counter++
	}

	return m
}

func getEncryptedCodes(s string, m map[string]int) []int {
	res := make([]int, 0, len(s))

	for _, v := range s {
		res = append(res, m[string(v)])
	}

	return res
}

func makeKeyLenEqualToCodesLen(strLen, key int) []int {
	strKey := strconv.Itoa(key)
	newKey := make([]int, 0, strLen)
	i := 0
	for len(newKey) < strLen {
		n, _ := strconv.Atoi(string(strKey[i]))
		newKey = append(newKey, n)

		if i+1 == len(strKey) {
			i = 0
		} else {
			i++
		}
	}
	return newKey
}

func Encode(str string, key int) []int {
	keyLetters := makeLettersKeyMap()
	encryptedCodes := getEncryptedCodes(str, keyLetters)
	newKey := makeKeyLenEqualToCodesLen(len(str), key)

	res := make([]int, 0, len(str))

	for i := 0; i < len(str); i++ {
		res = append(res, newKey[i]+encryptedCodes[i])
	}

	return res

}

func main() {
	fmt.Println(Encode("scout", 1939))
}
