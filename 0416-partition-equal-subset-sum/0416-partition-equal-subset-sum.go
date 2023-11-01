
func canPartition( nums []int) bool {
    totSum := 0
    n:=len(nums)
    for i := 0; i < n; i++ {
        totSum += nums[i]
    }

    if totSum%2 == 1 {
        return false
    } else {
        k := totSum / 2

        prev := make([]bool, k+1)
        for i := range prev {
            prev[i] = false
        }

        
        prev[0] = true

        if nums[0] <= k {
            prev[nums[0]] = true
        }

        for ind := 1; ind < n; ind++ {
            cur := make([]bool, k+1)
            for i := range cur {
                cur[i] = false
            }
            cur[0] = true

            for target := 1; target <= k; target++ {
                notTaken := prev[target]

                taken := false
                if nums[ind] <= target {
                    taken = prev[target-nums[ind]]
                }

                cur[target] = notTaken || taken
            }

            prev = cur
        }

       
        return prev[k]
    }
}
