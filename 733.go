package main

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	startingColor := image[sr][sc]
	if startingColor != color {
		dfs(image, sr, sc, color, startingColor)
	}
	return image
}

func dfs(image [][]int, sr int, sc int, color, startingColor int) {
	if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[0]) || image[sr][sc] != startingColor {
		return
	}
	image[sr][sc] = color
	dfs(image, sr-1, sc, color, startingColor)
	dfs(image, sr+1, sc, color, startingColor)
	dfs(image, sr, sc-1, color, startingColor)
	dfs(image, sr, sc+1, color, startingColor)
}
