func findDuplicate(nums []int) int {
    feq:=make(map[int]int)
    dup:=0
    for _,v:=range nums{
        feq[v]++
    }
    for key,v:=range feq{
        if v>1{
        dup=key
        }
    }
    return dup
}