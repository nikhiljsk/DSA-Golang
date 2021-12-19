/*
Rotate Image

You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.



Example 1:


Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]
Example 2:


Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
Example 3:

Input: matrix = [[1]]
Output: [[1]]
Example 4:

Input: matrix = [[1,2],[3,4]]
Output: [[3,1],[4,2]]


Constraints:

matrix.length == n
matrix[i].length == n
1 <= n <= 20
-1000 <= matrix[i][j] <= 1000
*/

// Refer this to understand this: https://dev.to/seanpgallivan/solution-rotate-image-cpp
// Basically i is the ring, j is the number of operations inside the ring
// Also, after each ring, the matrix is reduced by two dimensions
// To determinte indices, visualize that first in 00 moves towards right
// Similarly number in 40 moves up, number in 44 moves to left and number in 04 moves down
func rotate(matrix [][]int) {
	n := len(matrix)
	depth := n / 2
	for i := 0; i < depth; i++ {
		jlen := n - (2 * i) - 1
		opp := n - i - 1
		for j := 0; j < jlen; j++ {
			temp := matrix[i][i+j]
			matrix[i][i+j] = matrix[opp-j][i]
			matrix[opp-j][i] = matrix[opp][opp-j]
			matrix[opp][opp-j] = matrix[i+j][opp]
			matrix[i+j][opp] = temp
		}
	}
}

// Another solution. Easy to understand. Top, bottom, left and right pointers
func rotate(matrix [][]int) {
	n := len(matrix)
	l, r, t, b := 0, n-1, 0, n-1
	for i := 0; i < n/2; i++ {
		for j := 0; j < n-1-i*2; j++ {
			temp := matrix[t][l+j]
			matrix[t][l+j] = matrix[b-j][l]
			matrix[b-j][l] = matrix[b][r-j]
			matrix[b][r-j] = matrix[t+j][r]
			matrix[t+j][r] = temp
		}
		l++
		r--
		t++
		b--
	}
}

// Had to see the solution for this approach. Could get the intution behind the code later
// Took quite a while to get this around my head