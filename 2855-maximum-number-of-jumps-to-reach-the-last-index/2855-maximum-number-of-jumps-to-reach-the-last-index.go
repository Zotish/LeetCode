func maximumJumps(nums []int, target int) int {
    memo := make([]int, len(nums))
    for i := range memo {
        memo[i] = -1
    }
    return jump(nums, target, 0, memo)-1
}

func jump(nums []int, target int, currIndex int, memo []int) int {
    if currIndex == len(nums)-1 {
        return 1
    }

    if memo[currIndex] != -1 {
        return memo[currIndex]
    }

    maxJumps := 0
    for i := currIndex + 1; i < len(nums); i++ {
        diff := nums[i] - nums[currIndex]
        if -target <= diff && diff <= target {
            jumpsFromHere := jump(nums, target, i, memo)
            if jumpsFromHere != 0 {
                maxJumps = max(maxJumps, jumpsFromHere+1)
            }
        }
    }

    memo[currIndex] = maxJumps
    return maxJumps
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
