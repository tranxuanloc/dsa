package main

func findDiagonalOrder(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])
	var i, j int
	result := make([]int, m*n)
	up := true
	for k := range cap(result) {
		// fmt.Println(i, j, up)
		result[k] = mat[i][j]
		if up {
			if j+1 == n {
				i++
				up = false
			} else if i == 0 {
				j++
				up = false
			} else {
				i--
				j++
			}
		} else {
			if i+1 == m {
				j++
				up = true
			} else if j == 0 {
				i++
				up = true
			} else {
				i++
				j--
			}
		}
	}
	return result
}
