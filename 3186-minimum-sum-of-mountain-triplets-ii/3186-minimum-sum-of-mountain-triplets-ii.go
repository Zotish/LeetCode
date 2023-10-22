func minimumSum(nums []int) int {
    n := len(nums)

    
    leftMin := make([]int, n)
    rightMin := make([]int, n)


    leftMin[0] = nums[0]
    for i := 1; i < n; i++ {
        leftMin[i] = min(leftMin[i-1], nums[i])
    }

   
    rightMin[n-1] = nums[n-1]
    for i := n-2; i >= 0; i-- {
        rightMin[i] = min(rightMin[i+1], nums[i])
    }

    sum := int(1e9) 
    for j := 1; j < n-1; j++ {
        if nums[j] > leftMin[j] && nums[j] > rightMin[j] {
            sum = min(sum, nums[j]+leftMin[j]+rightMin[j])
        }
    }

    if sum == int(1e9) {  
        return -1
    }
    return sum
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
