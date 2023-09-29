//Brute force Approach
/*func subsetsWithDup(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums) // Sort the nums to handle duplicates.

	totalSubsets := 1 << n
	var result [][]int
	seen := make(map[string]bool)

	for i := 0; i < totalSubsets; i++ {
		var subset []int
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				subset = append(subset, nums[j])
			}
		}

		// Convert subset to string to use it as a key in the map.
		// This is a simple way to check for duplicates.
		key := fmt.Sprint(subset)

		if !seen[key] {
			result = append(result, subset)
			seen[key] = true
		}
	}

	return result
}
*/
// Recursive Approach
func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
   var ans [][]int
   ds:=make([]int,0) 
   findCombination(0,nums,&ans,ds)
   return ans
}
func findCombination(ind int,nums []int,ans *[][]int,ds []int){
           tmp := make([]int, len(ds))
		   copy(tmp, ds)
		   *ans = append(*ans, tmp)
   for i:=ind;i<len(nums);i++{
       if i!=ind && nums[i]==nums[i-1]{
           continue
       }
        ds=append(ds,nums[i])
        findCombination(i+1,nums,ans,ds)
        ds=ds[:len(ds)-1]
   }
}