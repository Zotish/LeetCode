func runningSum(nums []int) []int {
    sumArr:=make([]int,0)
    sum:=0
    for i:=0;i<len(nums);i++{
        sum+=nums[i]
        sumArr=append(sumArr,sum)
    }
    return sumArr
}