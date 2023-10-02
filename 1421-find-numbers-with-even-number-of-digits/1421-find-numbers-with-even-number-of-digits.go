func findNumbers(nums []int) int {
    count:=0
  for i:=0;i<len(nums);i++{
      getCount:=len(strconv.Itoa(nums[i]))
      if getCount%2==0{
          count++
      }
  }
  return count
}