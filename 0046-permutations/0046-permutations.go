func permute(nums []int) [][]int {
var ans [][]int
index:=0
ans=solve(nums, index, ans)
return ans
}
func solve(nums []int,index int,ans [][]int)[][]int  {
    if index>=len(nums){
        temp:= make([]int, len(nums))
        copy(temp,nums)
        ans=append(ans,temp)
        return ans
    }
    for j:=index;j<len(nums);j++{
        nums[index],nums[j]=nums[j],nums[index]
        ans=solve(nums,index+1,ans)
        nums[index],nums[j]=nums[j],nums[index]
    }
    return ans
}