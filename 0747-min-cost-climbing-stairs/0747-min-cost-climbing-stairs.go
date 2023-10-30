func minCostClimbingStairs(cost []int) int {
n:=len(cost)
dp:=make([]int,n+1)
  for i:=0;i<len(cost);i++{
      dp[i]=-1
  }
  dp[0]=cost[0]
  dp[1]=cost[1]
  for i:=2;i<n;i++{
     left:=cost[i]+dp[i-1]
     right:=cost[i]+dp[i-2]
     dp[i]=min(left,right)
  }
  return min(dp[n-1],dp[n-2])
}
func min(a,b int) int{
    if a>b{
        return b
    }
    return a
}