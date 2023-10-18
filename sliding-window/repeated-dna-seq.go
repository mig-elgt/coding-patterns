package slidingwindow

func findRepeatedSequences(s string, k int) Set {
	output := *NewSet()
	seqs := *NewSet()
	var start, end int
	for end < len(s) {
		if end-start == k-1 {
			seq := s[start : end+1]
			if seqs.Exists(seq) {
				if !output.Exists(seq) {
					output.Add(seq)
				}
			} else {
				seqs.Add(seq)
			}
			start++
		}
		end++
	}
	return output
}
