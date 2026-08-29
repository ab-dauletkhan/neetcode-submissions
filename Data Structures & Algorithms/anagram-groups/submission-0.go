func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, s := range strs {
		key := collect(s)
		m[key] = append(m[key], s)
	}

	return flatten(m)
}

func collect(s string) [26]int {
	var res [26]int

	for _, r := range s {
		res[r - 'a']++
	}

	return res
}

func flatten(m map[[26]int][]string) [][]string {
	var res [][]string

	for _, slc := range m {
		res = append(res, slc)
	}

	return res
}