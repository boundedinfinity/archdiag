package archdiag

func getStr(ss ...string) string {
	var f string

	for i := len(ss) - 1; i >= 0; i-- {
		c := ss[i]

		if c != "" {
			f = c
			break
		}
	}

	return f
}

func getInt(is ...int) int {
	var f int

	for i := len(is) - 1; i >= 0; i-- {
		c := is[i]

		if c != 0 {
			f = c
			break
		}
	}

	return f
}
