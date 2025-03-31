package main

func updateMatrix(mat [][]int) [][]int {
	r := len(mat)
	c := len(mat[0])
	queue := make([][2]int, 0)
	ans := make([][]int, r)
	for i := 0; i < r; i++ {
		ans[i] = make([]int, c)
		for j := 0; j < c; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			} else {
				ans[i][j] = -1
			}
		}
	}
	dirs := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		cr := cur[0]
		cc := cur[1]
		for _, dir := range dirs {
			nr := cr + dir[0]
			nc := cc + dir[1]
			if nr < 0 || nr >= r || nc < 0 || nc >= c || ans[nr][nc] != -1 {
				continue
			}
			queue = append(queue, [2]int{nr, nc})
			ans[nr][nc] = ans[cr][cc] + 1
		}
	}
	return ans
}
