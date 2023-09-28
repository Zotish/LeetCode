func reverseString(s []byte) {
    helper(s, 0, len(s)-1)
}

func helper(s []byte, i, j int) {
    if i >= j {
        return
    }
    s[i], s[j] = s[j], s[i]
    helper(s, i+1, j-1)
}