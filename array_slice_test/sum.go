package main

import "fmt"

func Sum(arr []int) int {
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))

	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// ":" meaning take one from the end of the array - slices a slice[low:high].
			tail := numbers[1:]
			fmt.Println(tail)
			sums = append(sums, Sum(tail))
		}

	}

	return sums
}
