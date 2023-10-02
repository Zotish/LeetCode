func countBits(n int) []int {
    arr:=make([]int,0)
    for i:=0;i<=n;i++{
        arr=append(arr,counter(i))
    }
    return arr
}
func counter(n int) int{
    count:=0
    for n>0{
        n &=n-1
        count++
    }
    return count
}