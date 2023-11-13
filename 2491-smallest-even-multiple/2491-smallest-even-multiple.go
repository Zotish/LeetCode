func smallestEvenMultiple(n int) int {
    var ans int 
    if n==1{
        ans=2
    }
    for i:=1;i<=n;i++{
       if n*i%2==0{
           ans=n*i
           break
       }
    }
    return ans
}