func newRoadSystem(roadRegister [][]bool) bool {

	totalRegisters := len(roadRegister)

	in := make([]int, totalRegisters)
	out := make([]int, totalRegisters)

	for k, v := range roadRegister {
		s := len(v)
		for i := 0; i < s; i++ {
			if v[i] == true {
				out[k] += 1
				in[i] += 1
			}
		}
	}

	same := true
	for i := 0; i < totalRegisters; i++ {
		if in[i] != out[i] {
			same = false
			break
		}
	}
	return same
}
