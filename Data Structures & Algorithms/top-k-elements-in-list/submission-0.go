func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	values := make([][]int, len(nums)+1)

	for num, idx := range counts {
		values[idx] = append(values[idx], num)
	}

	res := make([]int, 0, k)

	for i := len(values) - 1; i >= 0; i-- {
		if values[i] == nil {
			continue
		}

		for _, num := range values[i] {
			if len(res) < k {
				res = append(res, num)
			}
		}
	}

	return res
}
