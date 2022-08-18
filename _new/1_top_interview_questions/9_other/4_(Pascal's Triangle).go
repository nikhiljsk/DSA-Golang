// Approach 1 - Own
func generate(numRows int) [][]int {
	var res [][]int
	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}

	for i := 2; i <= numRows; i++ {
		res = append(res, generateNextRow(res[len(res)-1]))
	}
	return res
}

func generateNextRow(row []int) []int {
	var res []int
	res = append(res, row[0])
	for i := 0; i < len(row)-1; i++ {
		res = append(res, row[i]+row[i+1])
	}
	res = append(res, row[len(row)-1])
	return res
}

// Approach 2 - Solution
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 5M
// Time Implement - 15M