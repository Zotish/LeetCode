func countSubarrays(nums []int, k int64) int64 {
    l,r,count,sum:=0,0,0,0
    for ;r<len(nums);r++{
        sum+=nums[r]
        for {
         if int64(sum*(r-l+1))<k{
             break
         }
        sum-=nums[l]
        l++
        }
        count+=r-l+1
    }
    return int64(count)
}
