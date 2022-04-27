package nodedegree

import "fmt"

// Degree func
func Degree(nodes int, graph [][2]int, node int) (int, error) {
	ans := 0
	for _, edge := range graph {
		if edge[0] == node {
			ans++
		} else if edge[1] == node {
			ans++
		}
	}

	if ans == 0 {
		return 0, fmt.Errorf("node %d not found in the graph", node)
	}

	return ans, nil
}
