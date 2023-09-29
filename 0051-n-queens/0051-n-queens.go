func solveNQueens(n int) [][]string {
  var ans [][]string
  board := make([]string, n)
  s := strings.Repeat(".", n)
  for i:=0;i<n;i++{
      board[i]=s
  }
  leftrow := make([]int, n)
   upperDiagonal := make([]int, 2*n-1)
   lowerDiagonal := make([]int, 2*n-1)
  solve(0,board,&ans,leftrow,upperDiagonal,lowerDiagonal,n)
  return ans
}
func solve(col int,board []string,ans *[][]string,leftrow []int,upperDiagonal []int,lowerDiagonal []int,n int){
    if col==n{
        temp:=make([]string,len(board))
        copy(temp,board)
        *ans = append(*ans, temp)
        return
    }
    for row:=0;row<n;row++{
        if leftrow[row]==0 && upperDiagonal[row+col]==0 && lowerDiagonal[n-1+row-col]==0{
            r := []rune(board[row])
			r[col] = 'Q'
			board[row] = string(r)
            leftrow[row]=1
            upperDiagonal[row+col]=1
            lowerDiagonal[n-1+row-col]=1
            solve(col+1,board,ans,leftrow,upperDiagonal,lowerDiagonal,n)
            r[col] = '.'
			board[row] = string(r)
            leftrow[row]=0
            upperDiagonal[row+col]=0
            lowerDiagonal[n-1+row-col]=0   
        }
    }
}
