package main

import (
	"flag"
	"fmt"

	"github.com/tinygo-org/tinygo-site/doc-gen/cmd"
)

// main just calls the Execute function
func main() {

	createMachineDoc := flag.Bool("machine", false, "re-create machine package documentation for the given targets")
	updateBoardInterfaces := flag.Bool("interfaces", false, "update the Interfaces table for the given targets")
	updateBoardPins := flag.Bool("pins", false, "update the Pins table for the given targets")
	createHardwareMatrix := flag.Bool("matrix", false, "re-create the hardware matrix document")
	flag.Parse()

	// get the targets from the list of command line options
	targets := flag.Args()
	fmt.Printf("Targets: %+q\n", targets)

	cmd.Execute(targets, *createMachineDoc, *updateBoardInterfaces, *updateBoardPins, *createHardwareMatrix)
}
