func alternatingSums(a []int) []int {
	sum := []int{0, 0}
	for k, v := range a {
		if k%2 == 0 {
			sum[0] += v
		} else {
			sum[1] += v
		}
	}
	return sum
}
