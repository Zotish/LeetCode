func climbStairs(n int) int {
    if n==0{
        return 0
    }
    if n==1{
        return 1
    }
    
    prev:=1
    prev2:=1
    curr:=0
    for i:=1;i<n;i++{
        curr=prev+prev2
        prev2=prev
        prev=curr
    }
    return prev
}