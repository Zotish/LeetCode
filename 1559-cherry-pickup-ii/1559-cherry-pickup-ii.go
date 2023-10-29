func cherryPickup(grid [][]int) int {
    n := len(grid)
	m := len(grid[0])

	front := make([][]int, m)
	cur := make([][]int, m)
	for i := range front {
		front[i] = make([]int, m)
		cur[i] = make([]int, m)
	}

	for j1 := 0; j1 < m; j1++ {
		for j2 := 0; j2 < m; j2++ {
			if j1 == j2 {
				front[j1][j2] = grid[n-1][j1]
			} else {
				front[j1][j2] = grid[n-1][j1] + grid[n-1][j2]
			}
		}
	}

	for i := n - 2; i >= 0; i-- {
		for j1 := 0; j1 < m; j1++ {
			for j2 := 0; j2 < m; j2++ {
				maxi := -1 << 31 
				for di := -1; di <= 1; di++ {
					for dj := -1; dj <= 1; dj++ {
						var ans int
						if j1 == j2 {
							ans = grid[i][j1]
						} else {
							ans = grid[i][j1] + grid[i][j2]
						}

						if (j1+di < 0 || j1+di >= m) || (j2+dj < 0 || j2+dj >= m) {
							ans += -1e9
						} else {
							ans += front[j1+di][j2+dj]
						}

						if ans > maxi {
							maxi = ans
						}
					}
				}
				cur[j1][j2] = maxi
			}
		}

		for j1 := 0; j1 < m; j1++ {
			for j2 := 0; j2 < m; j2++ {
				front[j1][j2] = cur[j1][j2]
			}
		}
	}

	return front[0][m-1]
}