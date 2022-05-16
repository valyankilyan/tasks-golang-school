package warriors

// здесь просто реализуем поиск компонентов связности, где
// вершины -- это единички, а ребра находятся на единичках, которые стоят рядом

// код мне не нравится, но я не знаю, как сделать красивее

type Node struct {
	x, y int
}

func IsIndex(a Node, l int) bool {
	return a.x >= 0 && a.y >= 0 && a.x < l && a.y < l
}

func Dfs(matrix, was [][]bool, v Node) {
	was[v.x][v.y] = true
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			ni := Node{v.x + i, v.y + j}
			if IsIndex(ni, len(was)) && matrix[ni.x][ni.y] && !was[ni.x][ni.y] {
				Dfs(matrix, was, ni)
			}
		}
	}
}

func Count(image string) int {
	length := 0
	for image[length] != byte('\n') {
		length++
	}

	matrix, was := make([][]bool, length), make([][]bool, length)
	for i := 0; i < length; i++ {
		matrix[i], was[i] = make([]bool, length), make([]bool, length)
		for j := 0; j < length; j++ {
			matrix[i][j] = image[i*(length+1)+j] == byte('1')
		}
	}

	ans := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if !was[i][j] && matrix[i][j] {
				Dfs(matrix, was, Node{i, j})
				ans++
			}
		}
	}

	return ans
}
