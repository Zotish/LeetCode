func sortPeople(names []string, heights []int) []string {
   pairs := make([][2]interface{}, len(names))
	for i, n := range names {
		pairs[i] = [2]interface{}{n, heights[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1].(int) > pairs[j][1].(int)
	})

	sortedNames := make([]string, len(names))
	for i, pair := range pairs {
		sortedNames[i] = pair[0].(string)
	}

	return sortedNames 
}