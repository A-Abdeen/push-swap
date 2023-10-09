package pushswap

func RR(stack []int, which string) ([]int, bool, string) {
	var newStack []int
	for i := 0; i < len(stack); i++ {
		newStack = append(newStack, stack[i])
	}
	if len(stack) < 2 {
		return stack, false, "rr"
	}
	tmp := newStack[len(newStack)-1]
	for i := len(newStack) - 1; i > 0; i-- {
		newStack[i] = stack[i-1]
	}
	newStack[0] = tmp
	if Testlist(stack, newStack, which) {
		return newStack, true, "rr"
	} else {
		return stack, false, "rr"
	}
}
