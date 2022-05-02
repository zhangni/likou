package hot100

func searchMatrix(matrix [][]int, target int) bool {
	var m = len(matrix)
	var n = len(matrix[0])
	i := 0
	j := n - 1
	for i < m && j >= 0 {
		
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] == target {
			return true
		} else {
			i++
		}

	}
	return false /**/
}
