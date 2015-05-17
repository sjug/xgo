package pascal

func Triangle(n int) [][]int {
	pascal := make([][]int, n)
	for i := 0; i < n; i++ {
		pascal[i] = make([]int, i+1)
		pascal[i][0] = 1
		for j := 1; j < i; j++ {
			pascal[i][j] = pascal[i-1][j-1] + pascal[i-1][j]
		}
		pascal[i][i] = 1
	}
	return pascal
}
