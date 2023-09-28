func subsets(nums []int) [][]int {
	var res [][]int
	generateSubsets(0, []int{}, nums, &res)
	return res
}

func generateSubsets(index int, currentSubset []int, nums []int, res *[][]int) {
	if index == len(nums) {
		tmp := make([]int, len(currentSubset))
		copy(tmp, currentSubset)
		*res = append(*res, tmp)
		return
	}

	currentSubset = append(currentSubset, nums[index])
	generateSubsets(index+1, currentSubset, nums, res)

	currentSubset = currentSubset[:len(currentSubset)-1]
	generateSubsets(index+1, currentSubset, nums, res)
}