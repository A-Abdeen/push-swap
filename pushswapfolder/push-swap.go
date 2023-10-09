package main

import (
	"log"
	"os"
	pushswap "pushswap/functions"
	"fmt"
)

func main() {
	args := os.Args
	if len(args) == 2 {

	var stackA []int
	var argtest bool
	stackA, argtest = pushswap.MakeList(args[1]) // Check if file works correctly and arrange list into an array of int
	if !argtest{
		log.Fatal("Error")
	}
	pushswap.SortList(stackA)
	fmt.Print(pushswap.Combination)
	file, _:= os.Create("Output")
	_, _ = file.WriteString(pushswap.Combination)
	_, _ = file.WriteString("Final StackA:")
	_, _ = file.WriteString((pushswap.FinalStackA))
}
}

