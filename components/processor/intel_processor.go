package processor

import "fmt"

type intelProcessor struct {
}

func (p intelProcessor) Compute() {
	fmt.Println("Intel Computing...")
}

func NewIntelProcessor() *intelProcessor {
	return &intelProcessor{}
}
