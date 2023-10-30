func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    prev, prev2 := cost[0], cost[1]
    
    for i := 2; i < n; i++ {
        curri := cost[i] + min(prev, prev2)
        prev, prev2 = prev2, curri
    }
    
    return min(prev, prev2)
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
