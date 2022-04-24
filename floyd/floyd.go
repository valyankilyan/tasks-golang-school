package floyd

// Triangle makes a Floyd's triangle matrix with rows count.
func Triangle(rows int) [][]int {
	floyd := make([][]int, rows)
	l := 1
	for count := range floyd {
		floyd[count] = make([]int, count+1)
		for i := 0; i < count+1; i++ {
			floyd[count][i] = l
			l++
		}
	}
	return floyd
}
