func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	ans := math.MaxInt32
	left, sum := 0, 0

	for right := 0; right < n; right++ {
		sum += nums[right]
		for ;sum >= target;left++ {
            ans = int(math.Min(float64(ans),float64(right-left+1)))
			sum -= nums[left]
		}

	}
if ans == math.MaxInt32 {
		return 0
	}
	return ans
}