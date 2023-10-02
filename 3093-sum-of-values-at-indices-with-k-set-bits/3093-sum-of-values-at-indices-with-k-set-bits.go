func sumIndicesWithKSetBits(nums []int, k int) int {
    sum :=0
    for i:=0;i<len(nums);i++{
       if count(i)==k{
           sum+=nums[i]
       }
    }
    return sum
}
func count(n int) int{
    count1:=0
    for n>0{
        n &=n-1
        count1++
    }
    return count1
}
