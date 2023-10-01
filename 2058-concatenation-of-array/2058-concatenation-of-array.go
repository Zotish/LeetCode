func getConcatenation(nums []int) []int {
	arr := make([]int,len(nums)*2)
	for i,v := range nums {
		arr[i] = v
		arr[i+len(nums)] = v
	}
	return arr
}