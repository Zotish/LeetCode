
func minimumTotal(triangle [][]int) int {
    n:=len(triangle)
	front := make([]int, n)
	cur := make([]int, n)

	for j := 0; j < n; j++ {
		front[j] = triangle[n-1][j]
	}

	for i := n - 2; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			down := triangle[i][j] + front[j]
			diagonal := triangle[i][j] + front[j+1]
			cur[j] = int(math.Min(float64(down), float64(diagonal)))
		}
		copy(front, cur)
	}


	return front[0]
}