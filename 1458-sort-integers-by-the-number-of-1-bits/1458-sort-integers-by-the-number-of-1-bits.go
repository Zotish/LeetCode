func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		// Compare based on count of set bits
		x, y := count(arr[i]), count(arr[j])
		if x == y {
			// If counts are equal, compare based on actual value
			return arr[i] < arr[j]
		}
		return x < y
	})

	return arr
}

func count(n int) int {
	count := 0
	for n > 0 {
		n &= n - 1
		count++
	}
	return count
}