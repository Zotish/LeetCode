func countPairs(nums []int, target int) int {
    l:=0
    r:=len(nums)-1
    count:=0
    sort.Ints(nums)
    for l<r{
        sum:=nums[l]+nums[r]
        if sum<target{
            count+=r-l
            l++
        }else{
          r--
        }
    }
    return count
}