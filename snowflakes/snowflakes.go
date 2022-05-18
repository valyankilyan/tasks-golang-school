package snowflakes

//здесь edge -- это сторона треугольника
// in - цвет треугольника
// out - цвет снаружи треугольника
type edge struct {
	in  int
	out int
}

// выводим формулу характеристик и количества новых сторон треугольников
// исходя из своих абстрактных размышлений о геометрии

// дальше понимаем, что количество сторон любых характеристик всегда
//будет делиться на 3, столько же, сколько и сторон в треугольнике
// ну а так как треугольник может быть только одного цвета, мы можем с
// уверенностью с самого начала количество сторон поделить на 3

func OverlaidTriangles(n, m int) int {
	edges := make(map[edge]int)
	edges[edge{1, 0}] = 1
	for order := 1; order < n; order++ {
		new_edges := make(map[edge]int)
		for e, c := range edges {
			new_edges[e] += 2 * c
			new_edges[edge{e.in, e.in + 1}] += 1 * c
			new_edges[edge{1 + e.out, e.in + 1}] += 1 * c
			new_edges[edge{1 + e.out, e.out}] += 2 * c
		}
		edges = new_edges
	}

	ans := 0
	for e, c := range edges {
		if e.in == m {
			ans += c
		}
	}

	return ans
}
