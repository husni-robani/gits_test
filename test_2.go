package main

import (
	"fmt"
	"slices"
)

func main(){
	existing_scores := [][]int{{100, 100, 50, 40, 40, 20, 10 }, {100, 80, 80, 70}, {30, 20, 10}, {100, 90, 80}}
	gits_scores := [][]int{{5, 25, 50, 120}, {60, 70, 100}, {90, 100, 120}, {}}
	
	for i := 0; i < len(existing_scores); i++{
		fmt.Println("CASE ", i + 1)
		fmt.Println("---------------------------------")
		fmt.Println("Scores: ", existing_scores[i])
		fmt.Println("Gits Scores: ", gits_scores[i])
		result := solution( existing_scores[i], gits_scores[i])	
		fmt.Printf("Gits Ranks: %v\n\n", result)
	}
}

func solution(scores []int, gits_scores []int) []int {
	// remove duplicate scores
	for i := 0; ; {
		if i == len(scores) - 1 || len(scores) == 0 {
			break
		}
		if scores[i] == scores[i + 1]{
			scores = slices.Delete(scores, i, i + 1)
			continue
		}
		i ++
	}

	// get rank function
	getRank := func (score int) int {
		for i, v := range scores{
			if score >= v{
				return i + 1
			}
		}
		return len(scores) + 1
	}

	gits_ranks := make([]int, len(gits_scores))
	for i, v := range gits_scores{
		gits_ranks[i] = getRank(v)
	}

	return gits_ranks
}