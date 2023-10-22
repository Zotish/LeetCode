func minimumSum(nums []int) int {
    sum := int(1e9) 
    found := false

    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            for k := j + 1; k < len(nums); k++ {
                if nums[i] < nums[j] && nums[k] < nums[j] {
                    found = true
                    if sum > nums[i]+nums[j]+nums[k] {
                        sum = nums[i] + nums[j] + nums[k]
                    }
                }
            }
        }
    }

    if !found {
        return -1
    }

    return sum
}
