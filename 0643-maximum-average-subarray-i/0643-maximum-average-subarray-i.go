func findMaxAverage(nums []int, k int) float64 {
    sum := 0

    for i := 0; i < k; i++ {
        sum += nums[i]
    }

    maxSum := sum

    for i := k; i < len(nums); i++ {
        sum += nums[i] - nums[i-k]
        maxSum = max(maxSum, sum)
    }

    return float64(maxSum) / float64(k)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
