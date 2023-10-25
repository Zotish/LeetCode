const mod int = 1e9 * 2

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m, n := len(obstacleGrid), len(obstacleGrid[0])
    
    if obstacleGrid[0][0] == 1 {
        return 0
    }

    dp := make([]int, n)
    dp[0] = 1

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if obstacleGrid[i][j] == 1 {
                dp[j] = 0
                continue
            }
            
            if j-1 >= 0 {
                dp[j] = (dp[j] + dp[j-1]) % mod
            }
        }
    }
    return dp[n-1]
}
