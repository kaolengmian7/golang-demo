package bridge

import "fmt"

type printer interface {
	PrintFile()
}

type epson struct{}

func (p *epson) PrintFile() {
	fmt.Println("Printing by a EPSON printer")
}

type hp struct{}

func (p *hp) PrintFile() {
	fmt.Println("Printing by a HP printer")
}
