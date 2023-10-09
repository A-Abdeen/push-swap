package pushswap

func R(stack []int, which string) ([]int, bool, string) {
	var newStack []int
	for i := 0; i < len(stack); i++ {
		newStack = append(newStack, stack[i])
	}
	if len(stack) < 3 {
		return stack, false, "r"
	}
	tmp := newStack[0]
	for i := 0; i < len(newStack)-1; i++ {
		newStack[i] = newStack[i+1]
	}
	newStack[len(newStack)-1] = tmp
	if Testlist(stack, newStack, which) {
		return newStack, true, "r"
	} else {
		return stack, false, "r"
	}
}
