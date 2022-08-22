/*
Search a 2D Matrix

Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.


Example 1:


Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
Example 2:


Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false


Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104
*/

// Approach 1
func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

// Approach 2
func binarySearch(matrix []int, target int) bool {
	l := 0
	r := len(matrix) - 1
	for l <= r {
		mid := (l + r) / 2
		if matrix[mid] == target {
			return true
		} else if matrix[mid] > target {
			r--
		} else {
			l++
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	// Find the row the target could be in
	row := 0
	for i := 0; i < len(matrix); i++ {
		if matrix[i][len(matrix[0])-1] >= target {
			row = i
			break
		}
	}
	// Perform binary search on that row
	return binarySearch(matrix[row], target)
}

// Implemented