func maxOperations(nums []int, k int) int {
    sort.Ints(nums)
    l:=0
    r:=len(nums)-1
    count:=0
    for l<r{
        sum:=nums[l]+nums[r]
        if sum==k{
            l++
            r--
            count++
        }else if sum<k{
            l++
        }else{
            r--
        }
    }
    return count
}