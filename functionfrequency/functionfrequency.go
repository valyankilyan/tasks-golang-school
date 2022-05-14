package functionfrequency

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

import (
	"bufio"
	"fmt"
	"strings"
)

type Stack []int

func NewStack() Stack {
	q := Stack(make([]int, 0, 1))
	return q
}

func (q *Stack) Push(a ...int) {
	*q = append(*q, a...)
}

func (q *Stack) Pop() int {
	var ans int
	ans, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]
	return ans
}

// func IsQuote(a byte) bool {

// }

func FunctionFrequency(gocode []byte) []string {
	scanner := bufio.NewScanner(strings.NewReader(string(gocode)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// inQuotes, lastWord, leftBraces := false, []byte{}, NewStack()
	// strings.
	return []string{}
}
