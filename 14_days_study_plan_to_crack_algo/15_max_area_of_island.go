/*
Max Area of Island
You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

The area of an island is the number of cells with a value 1 in the island.

Return the maximum area of an island in grid. If there is no island, return 0.



Example 1:


Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
Output: 6
Explanation: The answer is not 11, because the island must be connected 4-directionally.
Example 2:

Input: grid = [[0,0,0,0,0,0,0,0]]
Output: 0


Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 50
grid[i][j] is either 0 or 1.
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Time - O(N)
// Iterate through each, once you find 1, get the biggest island for 1, and change it all to 2
// This way you don't reiterate the already calculated island section. So time complexity is order(Number of elements)
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	area := 0

	var count int
	var maxArea func(grid [][]int, i, j int) int
	maxArea = func(grid [][]int, i, j int) int {
		m, n := len(grid), len(grid[0])
		count++
		grid[i][j] = 2
		if i-1 >= 0 && grid[i-1][j] == 1 {
			maxArea(grid, i-1, j)
		}
		if i+1 < m && grid[i+1][j] == 1 {
			maxArea(grid, i+1, j)
		}
		if j-1 >= 0 && grid[i][j-1] == 1 {
			maxArea(grid, i, j-1)
		}
		if j+1 < n && grid[i][j+1] == 1 {
			maxArea(grid, i, j+1)
		}
		return count
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				area = max(area, maxArea(grid, i, j))
				count = 0
			}
		}
	}
	return area
}

// Implemented