/*
Reshape the Matrix
Add to List

Share
In MATLAB, there is a handy function called reshape which can reshape an m x n matrix into a new one with a different size r x c keeping its original data.

You are given an m x n matrix mat and two integers r and c representing the number of rows and the number of columns of the wanted reshaped matrix.

The reshaped matrix should be filled with all the elements of the original matrix in the same row-traversing order as they were.

If the reshape operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.



Example 1:


Input: mat = [[1,2],[3,4]], r = 1, c = 4
Output: [[1,2,3,4]]
Example 2:


Input: mat = [[1,2],[3,4]], r = 2, c = 4
Output: [[1,2],[3,4]]


Constraints:

m == mat.length
n == mat[i].length
1 <= m, n <= 100
-1000 <= mat[i][j] <= 1000
1 <= r, c <= 300
*/

// Time - O(m*n)
// Storage - O(m*n)
func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}
	// Store all the elements into a flat matrix
	flatMat := make([]int, 0, r*c)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			flatMat = append(flatMat, mat[i][j])
		}
	}
	// Create a new matrix of required shape
	retMat := make([][]int, r, r)
	for i := range retMat {
		retMat[i] = make([]int, c, c)
	}
	// Assign values
	k := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			retMat[i][j] = flatMat[k]
			k++
		}
	}
	return retMat
}

// Approach 2
// Time - O(m*n)
// Storage - O(1)
func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}
	// Create a new matrix of required shape
	retMat := make([][]int, r)
	r2 := 0
	// Assign values
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if len(retMat[r2]) == c {
				r2++
			}
			retMat[r2] = append(retMat[r2], mat[i][j])
		}
	}
	return retMat
}

/*
Testcases
[[1,2],[3,4]]
1
4
[[1,2],[3,4]]
2
4
[[1,2,3,21],[4,5,6,22],[7,8,9,23],[10,11,12,24]]
8
2
[[1,2,3,21],[4,5,6,22],[7,8,9,23],[10,11,12,24]]
1
2
[[1]]
1
1
[[1],[12],[13],[131]]
1
4
[[1],[12],[13],[131]]
4
1
*/

/*
// CONCEPT
// Len, Capacity, Slice

make(a, 0, 4)
0 4 []

make(a, 4, 4)
4 4 [[] [] [] []]

make(a, 4)
4 4 [[] [] [] []]

a := make([][]int, 4)
for i, _ := range a{
	a[i] = make([]int, 4, 4)
}
4 4 [[0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0]]

*/

// Implemented the first, had to see solution for second

