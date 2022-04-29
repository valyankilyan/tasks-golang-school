package functionfrequency

import (
	"fmt"
	"io"
)

type frequency map[string]int
type flow []byte

func (f *flow) ReadLine() ([]byte, error) {
	if len(*f) == 0 {
		return []byte{}, io.EOF
	}
	next_line := 0
	for (*f)[next_line] != '\n' {
		next_line++
	}
	ans := (*f)[:next_line]
	*f = (*f)[next_line+1:]
	return ans, nil
}

// func FuncFinder() return func (freq *frequency) (line []byte{
// 	in_quotes := false

// }

// func (freq *frequency) FindFunc(line []byte) {
// 	if bytes.HasPrefix(line, []byte("func ")) {
// 		return
// 	}
// 	func_names := make([]string, 0)
// 	last_left_

// }

func FunctionFrequency(gocode []byte) []string {
	var f flow = gocode
	var freq frequency = make(map[string]int)
	for {
		b, err := f.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(b))
	}
	return []string{}
}
