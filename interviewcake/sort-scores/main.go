package main

import "fmt"

// https://www.interviewcake.com/question/java/top-scores

// Map of scores and tally of occurrences
type Occurences map[int]int

func SortScores(scores []int, maxScore int) []int {
	results := make([]int, 0, len(scores))

	// Create map with maxScore space
	occurences := make(Occurences, maxScore)

	// Iterate through scores
	for _, score := range scores {
		// Fill map with scores and add tally
		occurences[score] += 1
	}

	// Iterate up to maxScore
	for i := 0; i <= maxScore; i++ {
		if occurences[i] <= 0 {
			continue
		}
		for j := occurences[i]; j > 0; j-- {
			// for tally many times, append value to results
			results = append(results, i)
		}

	}

	return results
}

func main() {
	scores := []int{50, 1, 4, 34, 1, 34, 4}
	results := SortScores(scores, 100)
	fmt.Println("input:", scores, "output:", results)
}
