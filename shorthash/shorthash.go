package shorthash

import (
	"math"
)

func recFind(cur string, dict []string, ans *[]string, left int) {
	if left == 0 {
		return
	}
	for _, s := range dict {
		*ans = append(*ans, cur+s)
		recFind(cur+s, dict, ans, left-1)
	}
}

func GenerateShortHashes(dictionary string, length int) []string {
	dictslice := make([]string, 0, len([]rune(dictionary)))
	for _, r := range dictionary {
		dictslice = append(dictslice, string(r))
	}

	anslength := int(math.Pow(float64(len(dictslice)), float64(length)))
	ans := make([]string, 0, anslength) // без перезаписи ))
	recFind("", dictslice, &ans, length)

	return ans
}
