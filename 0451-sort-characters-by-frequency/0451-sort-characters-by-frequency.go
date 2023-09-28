func frequencySort(s string) string {
	freq := make(map[rune]int)
	for _,v:=range s{
        freq[v]++
    }

	type Pair struct {
		Char  rune
		Count int
	}
	pairs := make([]Pair, 0, len(freq))
	for k, v := range freq {
		pairs = append(pairs, Pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	var result strings.Builder
	for _, p := range pairs {
		result.WriteString(strings.Repeat(string(p.Char), p.Count))
	}

	return result.String()
}