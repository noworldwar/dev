package main

import "fmt"

type human struct {
	name   string
	sexual string
}

func main() {
	terry := human{
		name:   "Terry",
		sexual: "Male",
	}

	terry.updateName("Retty")

	terry.updateNameByPointer("tttty")
	terry.print()
}

func (h human) print() {
	fmt.Printf("Current human is: %+v\n", h)
}

func (h human) updateName(newName string) {
	fmt.Printf("Before update: %+v\n", h)
	h.name = newName
	fmt.Printf("After update: %+v\n", h)
}

func (h *human) updateNameByPointer(newName string) {
	(*h).name = newName
}
