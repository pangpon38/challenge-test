package main

import (
	"fmt"
	"math"
)

func decodeSymbol(encoded string) string {

	n := len(encoded) + 1
	bestSeq := make([]int, n)
	bestSum := math.MaxInt32

	var recursiveFindBestResult func(index int, seq []int)
	recursiveFindBestResult = func(index int, seq []int) {

		if index == n {
			sum := 0
			for _, v := range seq {
				sum += v
			}

			if sum < bestSum {
				bestSum = sum
				copy(bestSeq, seq)
			}

			return
		}

		for i := 0; i <= 2; i++ {
			if index == 0 || (encoded[index-1] == 'L' && seq[index-1] > i) || (encoded[index-1] == 'R' && seq[index-1] < i) || (encoded[index-1] == '=' && seq[index-1] == i) {
				recursiveFindBestResult(index+1, append(seq, i))
			}
		}
	}

	recursiveFindBestResult(0, []int{})

	var result string
	for _, num := range bestSeq {
		result = fmt.Sprintf("%s%d", result, num)
	}

	return result
}

func main() {

	var encoded string

	fmt.Print("input => ")
	fmt.Scanln(&encoded)
	fmt.Println("output =>", decodeSymbol(encoded))
}
