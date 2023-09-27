func longestPrefix(s string) string {
	var max string
	for i := 1; i < len(s); i++ {
		if strings.HasPrefix(s, s[:i]) && strings.HasSuffix(s, s[:i]) {
			max = s[:i]
		}
	}
	return max
}