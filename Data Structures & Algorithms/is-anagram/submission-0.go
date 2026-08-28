func isAnagram(s string, t string) bool {
	return collect(s) == collect(t)
}

func collect(s string) [26]int {
	var res [26]int

	for _, r := range s {
		res[r - 'a']++
	}

	return res
}
