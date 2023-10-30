func minCostClimbingStairs(cost []int) int {
  dp:=make([]int,len(cost)+1)
  for i:=0;i<len(cost);i++{
      dp[i]=-1
  }
 return min(getMin(cost, len(cost)-1,dp), getMin(cost, len(cost)-2, dp))
}
func getMin(nums []int, n int,dp []int) int{
  if n<0{
      return 0
  }
  if dp[n]!=-1{
      return dp[n]
  }
  left:=nums[n]+getMin(nums,n-1,dp)
  right:=nums[n]+getMin(nums,n-2,dp)
  dp[n]=min(left,right)
  return dp[n]
}
func min(a,b int) int{
    if a>b{
        return b
    }
    return a
}