package bridge

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T) {

	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	macComputer := &mac{}
	macComputer.printer = hpPrinter
	macComputer.printer.PrintFile()
	fmt.Println()

	macComputer.printer = epsonPrinter
	macComputer.printer.PrintFile()
	fmt.Println()

	winComputer := &windows{}
	winComputer.printer = hpPrinter
	winComputer.printer.PrintFile()
	fmt.Println()

	winComputer.printer = epsonPrinter
	winComputer.printer.PrintFile()
	fmt.Println()
}
