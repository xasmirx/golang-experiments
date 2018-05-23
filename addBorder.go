func addBorder(picture []string) []string {

	var newPicture []string
	longest := 0

	for _, v := range picture {
		l := len(v)
		if l > longest {
			longest = l
		}
	}

	longest = longest + 2

	asterisk := make([]byte, longest)

	for i := 0; i < longest; i++ {
		asterisk[i] = '*'
	}

	newPicture = append(newPicture, string(asterisk))

	for _, v := range picture {
		l := len(v)
		tBytes := make([]byte, longest)

		for i := 0; i < longest; i++ {
			tBytes[i] = '*'
		}

		for j := 0; j < l; j++ {
			tBytes[j+1] = v[j]
		}

		newPicture = append(newPicture, string(tBytes))
	}

	newPicture = append(newPicture, string(asterisk))

	return newPicture
}
