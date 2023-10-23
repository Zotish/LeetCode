func rob(nums []int) int {
    prev:=nums[0]
    prev2:=0

    for i:=1;i<len(nums);i++{
        take:=nums[i]
        if i>1{
            take+=prev2
        }
        notTake:=0+prev
        curr:=max(take,notTake)
        prev2=prev
        prev=curr
    }
    return prev
}
func max(a,b int) int{
    if a>b{
        return a
    }
    return b
}