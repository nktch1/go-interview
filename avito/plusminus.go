package main

import (
	"fmt"
	"math"
)

// WIP

func prepareDigits(num int) []int {

	var (
		digits []int
		tmp    = num
	)

	for tmp > 0 {
		v := tmp % 10
		tmp /= 10

		digits = append(digits, v)
	}

	return digits
}

func maxMinus(variants []string) string {

	var (
		max    = math.MinInt32
		maxIdx = -1
	)

	comp := func(str string) int {
		res := 0
		template := rune('-')
		for _, el := range str {
			if el == template {
				res++
			}
		}

		return res
	}

	for idx, v := range variants {
		minusCount := comp(v)

		if minusCount > max {
			max = minusCount
			maxIdx = idx
		}
	}

	if maxIdx == -1 {
		return "not possible"
	}

	return variants[maxIdx]
}

func traverse(v int, digits []int, comb string, sum *int) {
	p := fmt.Sprintf("%s+", comb)
	n := fmt.Sprintf("%s-", comb)

	if len(digits) == 0 {
		if *sum+v == 0 {
			poss = append(poss, p)
		}
		if *sum-v == 0 {
			poss = append(poss, n)
		}

		return
	}

	//traverse(digits, p, v+*sum)
	//traverse(digits, n, v-*sum)
}

var poss = []string{}

func PlusMinus(num int) string {
	//digits := prepareDigits(num)

	//traverse(digits[0], digits[1:], "", 0)

	return maxMinus(poss)
}

func main() {
	// do not modify below here, readline is our function
	// that properly reads in the input for you
	//fmt.Println(PlusMinus(readline()))
}
