package pushswap

func P(stack1 []int, stack2 []int) ([]int, []int) {
	var stack1New []int
	stack2 = append(stack2, stack1[0])
	var newStack []int
	for i := 0; i < len(stack2); i++ {
		newStack = append(newStack, stack2[i])
	}
	tmp := newStack[len(newStack)-1]
	for i := len(newStack) - 1; i > 0; i-- {
		newStack[i] = stack2[i-1]
	}
	newStack[0] = tmp
	for i := 1; i < len(stack1); i++ {
		stack1New = append(stack1New, stack1[i])
	}
	return stack1New, newStack
}
