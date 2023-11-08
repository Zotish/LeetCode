func numSubarraysWithSum(nums []int, goal int) int {
    return atmost(nums,goal)-atmost(nums,goal-1)
}
func atmost(nums []int, goal int) int{
    r,l,sum,ans:=0,0,0,0
    for ;r<len(nums);r++{
        sum+=nums[r]
        for l<=r && sum >goal{
            sum-=nums[l]
            l++
        }
        ans+=r-l+1
    }
    return ans
}