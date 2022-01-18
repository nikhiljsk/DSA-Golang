/*
Flood Fill

An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.

You are also given three integers sr, sc, and newColor. You should perform a flood fill on the image starting from the pixel image[sr][sc].

To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color), and so on. Replace the color of all of the aforementioned pixels with newColor.

Return the modified image after performing the flood fill.



Example 1:


Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, newColor = 2
Output: [[2,2,2],[2,2,0],[2,0,1]]
Explanation: From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel), all pixels connected by a path of the same color as the starting pixel (i.e., the blue pixels) are colored with the new color.
Note the bottom corner is not colored 2, because it is not 4-directionally connected to the starting pixel.
Example 2:

Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, newColor = 2
Output: [[2,2,2],[2,2,2]]


Constraints:

m == image.length
n == image[i].length
1 <= m, n <= 50
0 <= image[i][j], newColor < 216
0 <= sr < m
0 <= sc < n
*/

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	r, c := len(image), len(image[0])
	color := image[sr][sc]
	var helper func(sr, sc int)
	if color == newColor { // This condition is necessary to stop gng into a loop
		return image
	}
	helper = func(sr, sc int) {
		if image[sr][sc] == color {
			image[sr][sc] = newColor
			if sr-1 >= 0 {
				helper(sr-1, sc)
			}
			if sr+1 < r {
				helper(sr+1, sc)
			}
			if sc-1 >= 0 {
				helper(sr, sc-1)
			}
			if sc+1 < c {
				helper(sr, sc+1)
			}
		}
	}
	helper(sr, sc)
	return image
}

/*
Testcase that goes into infinite loop
[[0,0,0],[0,1,1]]
1
1
1
*/

// Implemented. Had to see the solution to get infinite loop condition