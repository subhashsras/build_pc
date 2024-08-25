package components

import "fmt"

type Ram struct {}

func (r Ram) LoadFromDisk() {
	fmt.Println("Loading from disk...")
}

func (r Ram) SendToProcessor() {
	fmt.Println("Sending to processor...")
}

