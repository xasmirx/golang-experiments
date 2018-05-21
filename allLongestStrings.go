func allLongestStrings(inputArray []string) []string {
	best := 0
	var newArray []string

	for _, v := range inputArray {
		l := len([]rune(v))

		if l > best {
			best = l
		}
	}

	for _, v := range inputArray {
		l := len([]rune(v))
		if l >= best {
			newArray = append(newArray, v)
		}
	}

	fmt.Println(best)

	return newArray

}
