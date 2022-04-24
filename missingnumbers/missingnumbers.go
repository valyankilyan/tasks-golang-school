package missingnumbers

func Missing(numbers []int) []int {
	exist := make([]bool, len(numbers)+2)
	for _, n := range numbers {
		exist[n-1] = true
	}
	var ans []int
	for num, e := range exist {
		if !e {
			ans = append(ans, num+1)
		}
	}
	return ans
}
