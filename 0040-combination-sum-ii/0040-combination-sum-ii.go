func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
   var ans [][]int
   ds:=make([]int,0) 
   findCombination(0,target,candidates,&ans,ds)
   return ans
}
func findCombination(ind int,target int,candidates []int,ans *[][]int,ds []int){
        if target==0{
           tmp := make([]int, len(ds))
		   copy(tmp, ds)
		   *ans = append(*ans, tmp)
           return
        }
   for i:=ind;i<len(candidates);i++{
       if i>ind && candidates[i]==candidates[i-1]{
           continue
       }
        if candidates[ind]>target{
        break
    }
        ds=append(ds,candidates[i])
        findCombination(i+1,target-candidates[i],candidates,ans,ds)
        ds=ds[:len(ds)-1]
   }
}