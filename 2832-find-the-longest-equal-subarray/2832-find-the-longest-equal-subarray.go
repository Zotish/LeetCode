func longestEqualSubarray(nums []int, k int) int {
    feq:=make(map[int]int)
    l,r,ans:=0,0,0
    max1:=0
    for ;r<len(nums);r++{
            feq[nums[r]]++
            max1=max(max1,feq[nums[r]])
        for ;(r-l+1)-max1>k;l++{
            feq[nums[l]]--
        }
        ans=max(ans,max1)
    }
    return ans
}
func max(a,b int ) int{
    if a>b{
        return a
    }
    return b
}