package jaro

import (
	"strings"
)

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Distance(word1 string, word2 string) float64 {
	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)

	if word1 == word2 {
		return 1.0
	}

	max_dist := Max(len(word1), len(word2))/2 - 1
	m2 := make([]int, len(word2))
	match := 0.0

	for i := range word1 {
		for j := Max(0, i-max_dist); j < Min(len(word2), i+max_dist+1); j++ {
			if word1[i] == word2[j] && m2[j] == 0 {
				match++
				m2[j] = int(match)
				break
			}
		}
	}

	if match == 0 {
		return 0.0
	}

	t, order := 0.0, 0
	for _, a := range m2 {
		if a != 0 {
			order++
			if a != order {
				t++
			}
		}
	}

	return (match/float64(len(word1)) +
		match/float64(len(word2)) +
		(match-t/2.0)/match) / 3.0
}
