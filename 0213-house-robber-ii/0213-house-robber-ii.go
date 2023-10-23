func rob(nums []int) int {
    temp1:=make([]int,0)
    temp2:=make([]int,0)
    if len(nums)==1{
        return nums[0]
    }
    for i:=0;i<len(nums);i++{
        if i!=0{
            temp1=append(temp1,nums[i])
        }
        if i!=len(nums)-1{
            temp2=append(temp2,nums[i])
        }
    }
    return max(rob1(temp1),rob1(temp2))
}
func rob1(nums []int) int {
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