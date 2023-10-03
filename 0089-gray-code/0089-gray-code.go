func grayCode(n int) []int {
    var v []int
    for i:=0; i<1<<n; i++ {
        v =append(v, i^(i>>1));
    }
    return v;
}