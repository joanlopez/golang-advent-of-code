package not_quite_lisp

func FinalFloor(seq string) (floor int) {
	floor = 0
	for _, c := range seq {
		if c == '(' {
			floor++
		} else {
			floor--
		}
	}
	return
}

func EnterBasementStep(seq string) int {
	floor := 0
	for i, c := range seq {
		if c == '(' {
			floor++
		} else {
			floor--
		}

		if floor == -1 {
			return i+1
		}
	}
	return -1
}

