package main

import "fmt"

func main(){
	test_case := []int {7, 12, 0}

	for i, input := range test_case{
		fmt.Println("CASE ", i + 1)
		fmt.Println("---------------------------------")
		fmt.Println("Input: ", input)
		result := solution(input)
		fmt.Printf("Result: %v\n\n", result)
	}
}

func solution(n_sequence int) []int{
	result := make([]int, n_sequence)
	formula := func(x int) int{
		return ((x * (x + 1)) / 2) + 1
	}

	for i := 0; i < n_sequence; i++{
		result[i] = formula(i)
	}

	return result
}