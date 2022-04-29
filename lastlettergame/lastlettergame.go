package lastlettergame

var graph [][]int
var was []bool
var max_depth int

// Здесь показалось уместным указывать переменную, которую возвращаю,
// потому что по контексту сразу понятно, что делает дфс
func Dfs(depth, v int) (path []int) {
	was[v] = true
	// Ругается, что нельзя в дефере сделать ничего кроме вызова функции (как я понял)
	// Не знаю, насколько то, что я сделал - нормальная практика..
	defer func(v int) { was[v] = false }(v)
	for _, u := range graph[v] {
		if !was[u] {
			if p := Dfs(depth+1, u); len(p) != 0 {
				path = p
			}
		}
	}
	if len(path) != 0 || depth > max_depth { // некрасиво
		if depth > max_depth { // но не знаю как ещё
			max_depth = depth
		}
		return append(path, v)
	}

	return path
}

func Sequence(dic []string) []string {
	graph = make([][]int, len(dic))
	was = make([]bool, len(dic))
	for v, s1 := range dic {
		for u, s2 := range dic {
			if len(s1) != 0 && len(s2) != 0 && s1[len(s1)-1] == s2[0] {
				graph[v] = append(graph[v], u)
			}
		}
	}

	var max_path []int
	for i := range graph {
		if path := Dfs(1, i); len(path) != 0 {
			max_path = path
		}
	}

	var ans []string
	for i, j := 0, len(max_path)-1; i < j; i, j = i+1, j-1 {
		max_path[i], max_path[j] = max_path[j], max_path[i]
	}

	for _, ind := range max_path {
		ans = append(ans, dic[ind])
	}

	return ans
}
