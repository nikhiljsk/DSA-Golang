func any_exists(find map[string]bool, s1, s2, s3 string) bool {
	if _, ok := find[s1]; ok {
		return true
	}
	if _, ok := find[s2]; ok {
		return true
	}
	if _, ok := find[s3]; ok {
		return true
	}
	return false
}

func isValidSudoku(board [][]byte) bool {
	find := make(map[string]bool)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val == '.' {
				continue
			}
			row := fmt.Sprintf("Row:%v;Val:%v", i, val)
			col := fmt.Sprintf("Col:%v;Val:%v", j, val)
			box := fmt.Sprintf("Box:%v & %v;Val:%v", (i / 3), (j / 3), val)
			if any_exists(find, row, col, box) {
				return false
			}
			find[row] = true
			find[col] = true
			find[box] = true
		}
	}
	return true
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - No (Missed the i/3 & j/3 part differently, also the keywords row,col,box were missing)
// Solution - No
// Hints - No
// Time Approach - 5M
// Time Implement - 20M