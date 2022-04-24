package buildword

const INF int = 100000009

func Match(word, fragment string) int {
	if len(word) < len(fragment) {
		return 0
	}
	for i := range fragment {
		if word[i] != fragment[i] {
			return 0
		}
	}
	return len(fragment)
}

func Rec(word string, fragments []string) int {
	mn := INF
	if len(word) == 0 {
		return 0
	}
	for _, fragment := range fragments {
		if l := Match(word, fragment); l != 0 {
			if ans := Rec(word[l:], fragments); mn > ans {
				mn = ans
			}
		}
	}
	if mn != INF {
		mn++
	}
	return mn
}

func BuildWord(word string, fragments []string) int {
	ans := Rec(word, fragments)
	if ans == INF {
		return 0
	} else {
		return ans
	}
}
