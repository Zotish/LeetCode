
func maximumTripletValue(nums []int) int64 {
    n := len(nums)
    maxRes := 0
    maxSoFar := nums[0]
    
    for j := 1; j < n-1; j++ {
        if nums[j] < maxSoFar {
            for k := j + 1; k < n; k++ {
                val := (maxSoFar - nums[j]) * nums[k]
                if val > maxRes {
                    maxRes = val
                }
            }
        }
        if nums[j] > maxSoFar {
            maxSoFar = nums[j]
        }
    }
    return int64(maxRes)
}
