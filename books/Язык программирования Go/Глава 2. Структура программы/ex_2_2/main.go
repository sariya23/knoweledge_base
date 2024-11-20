package main

import (
	"bufio"
	"fmt"
	"main/tempconv"
	"os"
	"strconv"
)

type Info struct {
	From     tempconv.Celsius
	To       float64
	ToSymbol string
}

func main() {
	args := os.Args[1:]
	var data []tempconv.Celsius
	if len(args) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			v := scanner.Text()
			n, err := fetchData(scanner.Text())
			if err != nil {
				fmt.Printf("cannot parse value: %v\n", v)
				continue
			}
			data = append(data, n)
		}
	} else {
		for _, v := range args {
			n, err := fetchData(v)
			if err != nil {
				fmt.Printf("cannot parse value: %v\n", v)
				continue
			}
			data = append(data, n)
		}
	}
	res := convert(data)
	for _, v := range res {
		fmt.Printf("%s = %v%v\n", v.From, v.To, v.ToSymbol)
	}
}

func convert(t []tempconv.Celsius) []Info {
	res := make([]Info, 0, len(t)*2)
	for _, v := range t {
		res = append(res, Info{From: v, To: float64(v.CToF()), ToSymbol: "F"})
		res = append(res, Info{From: v, To: float64(v.CToK()), ToSymbol: "K"})
	}
	return res
}

func fetchData(v string) (tempconv.Celsius, error) {
	n, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, err
	}
	return tempconv.Celsius(n), nil
}
