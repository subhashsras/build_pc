package processor

import "fmt"

type amdProcessor struct {}

func (p amdProcessor) Compute() {
	fmt.Println("AMD Computing...")
}

func NewAMDProcessor()  *amdProcessor {
	return &amdProcessor{};
}