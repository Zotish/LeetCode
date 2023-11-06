func countNegatives(grid [][]int) int {
    m:=len(grid)
    n:=len(grid[0])
    count:=0
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if grid[i][j]<0{
                count++
            }
        }
    }
    return count
}