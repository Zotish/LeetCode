/*func permute(nums []int) [][]int {
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
}*/
func permute(nums []int) [][]int{
    var ans [][]int
    ds:=make([]int,0)
    feq:=make([]bool,len(nums))
    for i:=0;i<len(nums);i++{
        feq[i]=false
    }
    findPermute(ds,nums,&ans,feq)
    return ans
}
func findPermute(ds []int,nums []int,ans *[][]int,feq []bool){
    if len(ds)==len(nums){
        tmp := make([]int, len(ds))
		copy(tmp, ds)
		*ans = append(*ans, tmp)
        return
    }
   for i:=0;i<len(nums);i++{
      if !feq[i] {
          ds=append(ds,nums[i])
          feq[i]=true
          findPermute(ds,nums,ans,feq)
          feq[i]=false
          ds=ds[:len(ds)-1]
      }
   }
}