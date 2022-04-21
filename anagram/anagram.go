package anagram

import (
	"fmt"
	"strings"
)

func IsLetter(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func IsWord(word string) (ans bool) {
	for _, r := range word {
		ans = ans || IsLetter(r)
	}
	return ans
}

func CompareMasks(a, b []uint) bool {
	ans := true
	if len(a) != len(b) {
		fmt.Printf("len(a) = %d, len(b) = %d\n", len(a), len(b))
	}
	for i := range a {
		ans = ans && (a[i] == b[i])
		if !ans {
			return ans
		}
	}
	return ans
}

func FindMask(word string) []uint {
	word = strings.ToLower(word)
	ans := make([]uint, int('z')-int('a')+1)
	for _, letter := range word {
		if IsLetter(letter) {
			ans[int(letter)-int('a')]++
		}
	}
	return ans
}

func AnagramFinder(word string) func(cword string) bool {
	mask := FindMask(word)
	return func(cword string) bool {
		return CompareMasks(mask, FindMask(cword))
	}
}

func FindAnagrams(dictionary []string, word string) (result []string) {
	word = strings.ToLower(word)
	if !IsWord(word) {
		return
	}
	IsAnagram := AnagramFinder(word)
	for _, s := range dictionary {
		if strings.ToLower(s) != word && IsAnagram(s) {
			result = append(result, s)
		}
	}
	return
}
