func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    prev := make([]int, n) // Corrected initialization
    for i := 0; i < m; i++ {
        curr := make([]int, n) // Corrected initialization
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 {
                curr[j] = grid[i][j]
            } else {
                up := grid[i][j]
                if i > 0 {
                    up += prev[j]
                } else {
                    up += 1e9
                }
                left := grid[i][j]
                if j > 0 {
                    left += curr[j-1]
                } else {
                    left += 1e9
                }
                curr[j] = min(up, left)
            }
        }
        prev = curr
    }
    return prev[n-1]
}

func min(a, b int) int {
    if a < b { 
        return a
    }
    return b
}
