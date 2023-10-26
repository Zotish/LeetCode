func sortArrayByParity(nums []int) []int {
    arr1:=make([]int,0)
    arr2:=make([]int,0)
    arr3:=make([]int,0)
    for i:=0;i<len(nums);i++{
        if nums[i]%2==0 && nums[i]>=0{
            arr1=append(arr1,nums[i])
        }else {
            arr2=append(arr2,nums[i])
        }
    }
 arr3=append(arr1,arr2...)
 return arr3
}