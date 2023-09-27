func findAnagrams(s string, p string) []int {
    ans:=[]int{}
    var Temp [26]int
    var Hash [26]int
  for i:=0;i<len(p);i++{
      Hash[p[i]-'a']++
  }
  l:=0
  r:=0
  for ;r<len(s);r++{
      if Temp==Hash{
          ans=append(ans,l)
      }
        Temp[s[r]-'a']++
      for ;r-l+1>len(p);l++{
          Temp[s[l]-'a']--
      }
  }
  if Temp==Hash{
      ans=append(ans,l)
  }
  return ans
}