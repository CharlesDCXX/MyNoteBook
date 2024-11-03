package main

func numberOfRightTriangles(grid [][]int) int64 {
	var hang, lie []int
	for i := 0; i < len(grid); i++ {
		hang_sum := 0
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				hang_sum += 1
			}
		}

		hang = append(hang, hang_sum)
	}
	for i := 0; i < len(grid[0]); i++ {
		lie_sum := 0
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == 1 {
				lie_sum += 1
			}
		}
		lie = append(lie, lie_sum)

	}

	result := 0
	for i := 0; i < len(hang); i++ {
		for j := 0; j < len(lie); j++ {
			if hang[i] > 1 && lie[j] > 1 && grid[i][j] == 1 {
				result += (hang[i] - 1) * (lie[j] - 1)
			}
		}
	}
	return int64(result)
}
func main4() {
	grid := [][]int{{1, 0, 0, 0}, {0, 1, 0, 1}, {1, 0, 0, 0}}
	numberOfRightTriangles(grid)

	// 1 0 0 0
	// 0 1 0 1
	// 1 0 0 0
}
