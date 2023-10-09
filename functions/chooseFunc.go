package pushswap

func ChooseFunc(stackA []int, stackB []int, f func([]int, string) ([]int, bool, string)) ([]int, []int) {
	stackA, testA, fName := f(stackA, "A")
	stackB, testB, _ := f(stackB, "B")
	if testA {// depending on the func choosen 
		if testB {
			if fName == "rr" {
				fName = "rrr"
			} else {
				fName += fName
			}
			Combination = Combination + fName + "\n"
		} else {
			Combination = Combination + fName + "a" + "\n"
		}
	} else {
		if testB {
			Combination = Combination + fName + "b" + "\n"
		}
	}
	return stackA, stackB
}
