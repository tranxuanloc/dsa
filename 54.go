package main

var dirs [4][2]int = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	result := make([]int, 0, m*n)
	helper54(0, m, 0, n, matrix, &result)
	return result
}

func helper54(sr, er, sc, ec int, matrix [][]int, result *[]int) {
	if sr > er || sc > ec {
		return
	}
	i, j := sr, sc
	d := 0
	// fmt.Println(sr, er, sc, ec)
	for i >= sr && i < er && j >= sc && j < ec {
		// fmt.Println(d, i, j)
		*result = append(*result, matrix[i][j])
		i += dirs[d][0]
		j += dirs[d][1]
		if d == 0 && j == ec {
			// down
			j--
			i++
			d++
		} else if d == 1 && i == er {
			// left
			i--
			j--
			d++
		} else if d == 2 && j < sc {
			// up
			i--
			j++
			d++
		}
		if i == sr && j == sc {
			break
		}
	}
	helper54(sr+1, er-1, sc+1, ec-1, matrix, result)
}
