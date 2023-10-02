func findTheDifference(s string, t string) byte {
    var res byte = 0

		for i:=0;i<len(t);i++ {
			if i < len(s) {
				res ^= s[i]
			}
			res ^=t[i] 
		}
		return res
}