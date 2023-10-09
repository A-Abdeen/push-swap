package pushswap

import "strconv"

var Combination string
var FinalStackA string

func SortList(stackA []int) string {
	var stackB []int
	switchP := 1
	for i := 0; i >= 0; i++ {
		oldStackA := stackA
		oldStackB := stackB
		stackA, stackB = ChooseFunc(stackA, stackB, R)
		stackA, stackB = ChooseFunc(stackA, stackB, S)
		stackA, stackB = ChooseFunc(stackA, stackB, RR)
		if TestFinalList(stackA, "A") && TestFinalList(stackB, "B") {// if both stacks are arranged break
			break
		}
		if Testlist(oldStackA, stackA, "A") || Testlist(oldStackB, stackB, "B") {// check if the stack changed if so go through the functions again
			i--
			continue
		} else {
			if switchP%2 == 0 {
				if len(stackB) == 1 { // if stackB is empty switch P function from a to b
					switchP++
				}
				stackB, stackA = P(stackB, stackA)
				Combination = Combination + "pa\n"
			} else {
				if len(stackA) == 1 {
					switchP++
				}
				stackA, stackB = P(stackA, stackB)
				Combination = Combination + "pb\n"
			}
		}
	}
	for i := 0; i >= 0; i++ {
		var testA bool
		if len(stackB) == 0 {// return all the values from stackb to stacka
			break
		}
		stackB, stackA = P(stackB, stackA)
		Combination = Combination + "pa\n"
		stackA, testA, _ = S(stackA, "A")
		if testA {
			Combination = Combination + "sa\n"
		}
	}
	for i := 0; i < len(stackA); i++ {
		FinalStackA = (FinalStackA + " " + strconv.Itoa(stackA[i]))
	}
	return Combination
}
