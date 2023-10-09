package main
import (
	"bufio"
	"fmt"
	"os"
	pushswap "pushswap/functions"
)
func main() {
	args := os.Args
	if len(args) == 2 {
stackA, argtest := pushswap.MakeList(args[1])
if !argtest{
	fmt.Println("Error")
}else{
	r := bufio.NewReader(os.Stdin)
    name, _ := r.ReadString(' ')
	checker := pushswap.Checker(stackA, string(name))
	fmt.Println(checker)
}
}
}