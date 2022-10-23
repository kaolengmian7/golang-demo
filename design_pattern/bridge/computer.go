package bridge

import "fmt"

type computer interface {
	Run()
}

type mac struct {
	printer printer
}

func (m *mac) Run() {
	fmt.Println("mac run")
}

type windows struct {
	printer printer
}

func (w *windows) Run() {
	fmt.Println("windows run")
}
