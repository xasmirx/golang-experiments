func candies(n int, m int) int {

	total := 0

	if n > m {
		return 0
	}

	for total <= m {
		total += n
	}

	if total > m {
		total -= n
	}

	return total

}
