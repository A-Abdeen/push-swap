package pushswap
func MakeList(argument string) ([]int, bool) {
	var stackA []int
	for i := 0; i < len(argument); i++ {
		if (argument[i] != 32) && ((argument[i] < 48) || (argument[i] > 59)) { // check if argument omly contains integers
			return stackA, false
		}
		if argument[i] != 32 {
			numb := int(argument[i] - 48)
			for j := i + 1; j < len(argument); j++ { // if the number is greater than 9
				if argument[j] == 32 {
					i = j
					break
				} else {
					numb = (numb * 10) + int(argument[j]-48)
					i = j
				}
			}
			stackA = append(stackA, numb)
		}
	}
	for l := 0; l < len(stackA); l++ { // check if value repeats
		for k := 0; k < len(stackA); k++ {
			if stackA[l] == stackA[k] && l != k {
				return stackA, false
			}
		}
	}
	return stackA, true
}
