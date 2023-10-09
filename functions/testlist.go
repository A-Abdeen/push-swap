package pushswap

// import "fmt"
func Testlist(stackA []int, stackANew []int, which string) bool {
	countOldA := 0
	countNewA := 0
	if which == "A" {
		for i := 0; i < len(stackA); i++ {
			for j := i + 1; j < len(stackA); j++ {// check list whenever it is ascending increase counter
				if stackA[i] < stackA[j] {
					countOldA++
				}
				if stackANew[i] < stackANew[j] {
					countNewA++
				}
			}
		}
	} else {
		for i := 0; i < len(stackA); i++ {// check list whenever it is descending increase counter
			for j := i + 1; j < len(stackA); j++ {
				if stackA[i] > stackA[j] {
					countOldA++
				}
				if stackANew[i] > stackANew[j] {
					countNewA++
				}
			}
		}
	}
	if countNewA > countOldA {
		return true
	} else {
		return false
	}
}
