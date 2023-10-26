func threeSum(nums []int) [][]int {
  var ans [][]int
  n:=len(nums)
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {

		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				temp := []int{nums[i], nums[j], nums[k]}
				ans = append(ans, temp)
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return ans
}
