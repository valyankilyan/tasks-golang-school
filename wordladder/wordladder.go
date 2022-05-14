package wordladder

// запускаем обход в ширину, и прямо там ищем ребра.
// асимптотика бфс с заранее известными ребрами n + m,
// а в нашем случае это будет n^2, но!!! не в самом худшем
// случае мы получим что-то меньшее, а с предподсчетом ребер
// мы всегда будем получать строгий n^2, поэтому я предпочел использовать
// первый вариант

type Queue []int

func NewQueue() Queue {
	q := Queue(make([]int, 0, 1))
	return q
}

func (q *Queue) Push(a ...int) {
	*q = append(*q, a...)
}

func (q *Queue) Pop() int {
	var ans int
	ans, *q = (*q)[0], (*q)[1:]
	return ans
}

func IsEdge(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	diff := 0
	for i := 0; i < len(a) && diff <= 1; i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff == 1
}

func FindEdges(from int, dic []string, depth []int) []int {
	edges := make([]int, 0)
	for i := range dic {
		if depth[i] == 0 && IsEdge(dic[from], dic[i]) {
			depth[i] = depth[from] + 1
			edges = append(edges, i)
		}
	}
	return edges
}

func WordLadder(from string, to string, dic []string) int {
	if from == to {
		return 1
	}
	depth := make([]int, len(dic))
	start, end := -1, -1
	q := NewQueue()

	for i := range depth {
		if dic[i] == from {
			depth[i] = 1
			q.Push(i)
			start = i
		} else if dic[i] == to {
			end = i
		}
	}

	if start == -1 {
		q.Push(len(dic))
		dic = append(dic, from)
		depth = append(depth, 1)
	}
	if end == -1 {
		dic = append(dic, to)
		depth = append(depth, 0)
	}

	for len(q) > 0 {
		cur := q.Pop()
		edges := FindEdges(cur, dic, depth)
		for _, e := range edges {
			if e == end {
				return depth[e]
			}
		}
		q.Push(edges...)
	}

	return 0
}
