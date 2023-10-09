package pushswap

func S(stack []int, which string) ([]int, bool, string) {
	var newStack []int
	for i := 0; i < len(stack); i++ {
		newStack = append(newStack, stack[i])
	}
	if len(stack) < 3 {
		return stack, false, "s"
	}

	newStack[0], newStack[1] = newStack[1], newStack[0]
	if Testlist(stack, newStack, which) {
		return newStack, true, "s"
	} else {
		return stack, false, "s"
	}
}
