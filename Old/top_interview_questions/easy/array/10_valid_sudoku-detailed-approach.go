/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.


Example 1:


Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true
Example 2:

Input: board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.


Constraints:

board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'.
*/
type Pair struct {
	i, j int
}

func newMap() map[byte]bool {
	return make(map[byte]bool)
}

func getBoxes(k, l int) []Pair {
	sub := []Pair{
		{0 + k, 0 + l}, {0 + k, 1 + l}, {0 + k, 2 + l},
		{1 + k, 0 + l}, {1 + k, 1 + l}, {1 + k, 2 + l},
		{2 + k, 0 + l}, {2 + k, 1 + l}, {2 + k, 2 + l},
	}
	return sub
}

func checkBoxes(board [][]byte, boxes []Pair) bool {
	temp := newMap()
	// fmt.Println(boxes)
	for _, obj := range boxes {
		// fmt.Println(" --    SB", obj)
		if _, ok := temp[board[obj.i][obj.j]]; ok {
			return false
		}
		if board[obj.i][obj.j] != byte('.') {
			// fmt.Println(" --    SB - Goes in:", board[obj.i][obj.j])
			temp[board[obj.i][obj.j]] = true
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	// Check rows & columns
	temp := newMap()
	temp2 := newMap()
	for i := 0; i < 9; i++ {
		temp = newMap()
		temp2 = newMap()
		for j := 0; j < 9; j++ {
			if _, ok := temp[board[i][j]]; ok {
				return false
			}
			if board[i][j] != byte('.') {
				temp[board[i][j]] = true
			}
			if _, ok := temp2[board[j][i]]; ok {
				return false
			}
			if board[j][i] != byte('.') {
				temp2[board[j][i]] = true
			}
		}
	}

	// Check sub-boxes
	x := true
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			boxes := getBoxes(i, j)
			x = x && checkBoxes(board, boxes)
			if !x {
				return x
			}
		}
	}

	return x

}

// Able to crack the approach and implemented