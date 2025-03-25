package main

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 49 {
				count++
				queue := [][2]int{
					{i, j},
				}
				for len(queue) > 0 {
					cur := queue[0]
					queue = queue[1:]
					for _, dir := range directions {
						r := cur[0] + dir[0]
						c := cur[1] + dir[1]
						if r < 0 || c < 0 || r == m || c == n || grid[r][c] != 49 {
							continue
						}
						grid[r][c] = 48
						queue = append(queue, [2]int{r, c})
					}
				}
			}
		}
	}
	return count
}
