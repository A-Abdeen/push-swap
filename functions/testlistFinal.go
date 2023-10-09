package pushswap

func TestFinalList(stackA []int, which string) bool {
	if which == "A" {
		for i := 0; i < len(stackA)-1; i++ {// check if stackA is ascending
			if stackA[i] < stackA[i+1] {
				continue
			} else {
				return false
			}
		}
	} else {
		for i := 0; i < len(stackA)-1; i++ {// check if stackB is descending
			if stackA[i] > stackA[i+1] {
				continue
			} else {
				return false
			}
		}
	}
	return true
}
