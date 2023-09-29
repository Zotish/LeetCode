func combinationSum(candidates []int, target int) [][]int {
   var ans [][]int
   ds:=make([]int,0) 
   findCombination(0,target,candidates,&ans,ds)
   return ans
}
func findCombination(ind int,target int,candidates []int,ans *[][]int,ds []int){
    if ind==len(candidates){
        if target==0{
           tmp := make([]int, len(ds))
		   copy(tmp, ds)
		   *ans = append(*ans, tmp)
        }
        return 
    }
    if candidates[ind]<=target{
        ds=append(ds,candidates[ind])
        findCombination(ind,target-candidates[ind],candidates,ans,ds)
        ds=ds[:len(ds)-1]
    }
    findCombination(ind+1,target,candidates,ans,ds)
}