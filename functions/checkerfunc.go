package pushswap

import ("log")

// import "fmt"

func Checker(stackA []int, command string) string {
	var commandArray []string
	var stackB []int
	count := 0
	lengthOfStack := len(stackA)
	for i := 0; i < len(command)-1; i++ {
		if command[i] != 10 && (command[i]<97 || command[i]>122){//if there is anything other than an integer print Error
			log.Fatal("Error")
		}
		if command[i] == 10 || command[i+1] == 10 {
			continue
		}
		if command[i+2] != 10 {
			commandArray = append(commandArray, (string(command[i]) + string(command[i+1]) + (string(command[i+2]))))
		} else {
		commandArray = append(commandArray, (string(command[i]) + string(command[i+1])))// append the data into an array
	}}
	for i := 0; i < len(commandArray); i++ {
		count++
		switch {
		case commandArray[i] == "pb":
			stackA, stackB = P(stackA, stackB)
		case commandArray[i] == "pa":
			stackB, stackA = P(stackB, stackA)
		case commandArray[i] == "sa":
			stackA, _, _ = S(stackA, "A")
		case commandArray[i] == "sb":
			stackB, _, _ = S(stackB, "B")
		case commandArray[i] == "ss":
			stackA, _, _ = S(stackA, "A")
			stackB, _, _ = S(stackB, "B")
		case commandArray[i] == "ra":
			stackA, _, _ = R(stackA, "A")
		case commandArray[i] == "rb":
			stackB, _, _ = R(stackB, "B")
		case commandArray[i] == "rr":
			stackA, _, _ = R(stackA, "A")
			stackB, _, _ = R(stackB, "B")
		case commandArray[i] == "rra":
			stackA, _, _ = RR(stackA, "A")
		case commandArray[i] == "rrb":
			stackB, _, _ = RR(stackB, "B")
		case commandArray[i] == "rrr":
			stackA, _, _ = RR(stackA, "A")
			stackB, _, _ = RR(stackB, "B")
		}
	}
	if TestFinalList(stackA, "A") && len(stackA) == lengthOfStack {// tests if stackA is sorted
		return "OK"
	}
	return "KO"
}
