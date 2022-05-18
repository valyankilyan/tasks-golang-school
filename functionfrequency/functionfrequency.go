package functionfrequency

import (
	"sort"
)

// regex howabout
// - find all positions of quotes
// - find all positions of structures like "[a-zA-Z.]("
// - then somehow determine if it was an actual call (doesn't have func before)
// - but there is lambda func which looks like func () {} or smth
// - well we could what?? okay

// scanner??
// - remove text in quotes
// - remove func*()
// - find all the [a-zA-Z.]( and remove (

type Stack []int

func NewStack() Stack {
	s := Stack(make([]int, 0, 1))
	return s
}

func (s *Stack) Push(a ...int) {
	*s = append(*s, a...)
}

func (s *Stack) Pop() int {
	var ans int
	ans, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return ans
}

func isUnusedStart(gocode []byte, ind int) bool {
	return gocode[ind] == '"' ||
		gocode[ind] == '\'' ||
		gocode[ind] == '`' ||
		gocode[ind] == '/' && ind > 0 && gocode[ind-1] == '/' ||
		gocode[ind] == '*' && ind > 0 && gocode[ind-1] == '/'
}

func isUnusedEnd(gocode []byte, ind int, unused *byte) bool {
	switch *unused {
	case '"':
		return gocode[ind] == '"'
	case '\'':
		return gocode[ind] == '\''
	case '`':
		return gocode[ind] == '`'
	case '/':
		return gocode[ind] == '\n'
	case '*':
		return gocode[ind-1] == '*' && gocode[ind] == '/'
	}
	return false
}

func InUnused(unused *byte, ind int, gocode []byte) bool {
	if *unused == 0 && isUnusedStart(gocode, ind) {
		*unused = gocode[ind]
		return true
	} else if *unused != 0 {
		if isUnusedEnd(gocode, ind, unused) {
			*unused = 0
		}
		return true
	}
	return false
}

func InDeclaration(lastWord []byte, declaration *bool, cur byte) bool {
	if !*declaration && string(lastWord) == "func" {
		*declaration = true
	} else if *declaration && cur == ')' {
		*declaration = false
	}
	return *declaration
}
func IsFunction(lastWord []byte, cur byte) bool {
	if len(lastWord) != 0 &&
		lastWord[len(lastWord)-1] != '.' &&
		cur == '(' {
		return true
	} else {
		return false
	}
}

func IsWord(l byte) bool {
	ans := l >= byte('a') && l <= byte('z')
	ans = ans || (l >= byte('A') && l <= byte('Z'))
	ans = ans || (l >= byte('0') && l <= byte('9'))
	return ans || l == byte('.') || l == byte('_')
}

func FunctionFrequency(gocode []byte) []string {
	unused, declaration, lastWordInd := byte(0), true, 0
	freq := make(map[string]int)

	for i, g := range gocode {
		if InUnused(&unused, i, gocode) {
			continue
		}
		if InDeclaration(gocode[lastWordInd:i], &declaration, g) {
			continue
		}
		if IsFunction(gocode[lastWordInd:i], g) {
			freq[string(gocode[lastWordInd:i])]++
		}

		if !IsWord(g) {
			lastWordInd = i + 1
		}
	}

	freqlist := make(PairList, len(freq))
	i := 0
	for k, v := range freq {
		freqlist[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(freqlist))
	var ans []string
	for i := 0; i < len(freqlist) && i < 3; i++ {
		ans = append(ans, freqlist[i].key)
	}

	return ans
}

type Pair struct {
	key   string
	value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].value < p[j].value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
