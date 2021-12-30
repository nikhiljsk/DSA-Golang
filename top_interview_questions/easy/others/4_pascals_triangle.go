/*
Pascal's Triangle

Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:




Example 1:

Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
Example 2:

Input: numRows = 1
Output: [[1]]


Constraints:

1 <= numRows <= 30
*/

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}

	for i := 1; i < numRows; i++ {
		res = append(res, genNext(res[i-1]))
	}

	return res
}

func genNext(arr []int) []int {
	newArr := make([]int, len(arr)+1)
	j := 1
	newArr[0] = 1
	for i := 0; i < len(arr)-1; i++ {
		newArr[j] = arr[i] + arr[i+1]
		j++
	}
	newArr[len(arr)] = 1
	return newArr
}

// Able to get the approach, but for solution had to see that loop is used to call another function
// Most implemented on my own.