package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/tinygo-org/tinygo-site/doc-gen/internal/boarddoc"
	"github.com/tinygo-org/tinygo-site/doc-gen/internal/hardwarematrix"
	"github.com/tinygo-org/tinygo-site/doc-gen/internal/machinepackagedoc"
	"github.com/tinygo-org/tinygo-site/doc-gen/internal/targetinfo"
)

var controllerDocBaseDir = filepath.Join("..", "content", "docs", "reference", "microcontrollers")

type documentStatus int

const (
	documentFound documentStatus = iota + 1
	missingDocOk
	missingDocButShouldHave
	missingDoc
	canProcessed
	finishedWithoutTargetOk
	finishedOk
)

type collection struct {
	target            string
	controllerDocPath string
	docStatus         documentStatus
	errors            []error
	matrixRow         *hardwarematrix.RowInfo
}

// this targets are virtual targets or common targets (chips) used as reference (inherits) in documented targets
var targetsKnownToSkip = map[string]struct{}{"atmega1284p": {}, "attiny1616": {}, "clue-alpha": {}, "cortex-m-qemu": {},
	"esp32": {}, "esp32c3": {}, "esp8266": {}, "riscv-qemu": {}, "simavr": {}, "wasi": {}, "wasm": {}}

// this targets are known to skip since last run, but are able to be documented
// TODO: validate this list
var targetsKnownToSkipAndShouldBeDocumented = map[string]struct{}{"bluemicro840": {}, "esp32-c3-devkit-rust-1": {},
	"feather-m0-express": {}}

func Execute(targets []string, createMachineDoc, updateBoardInterfaces, updateBoardPins, createHardwareMatrix bool) {
	// collect
	collector := collectFiles(targets)

	// update documents
	for target, co := range collector {
		if co.docStatus != canProcessed && co.docStatus != documentFound {
			continue
		}

		if co.docStatus == canProcessed || createHardwareMatrix {
			fmt.Printf("Processing '%s', please be patient\n", target)
		}

		if co.docStatus == canProcessed {
			if createMachineDoc || updateBoardInterfaces || updateBoardPins {
				fmt.Println(" Generating Board documentation...")
				machineDocPath := filepath.Join(controllerDocBaseDir, "machine", target+".md")
				if err := updateMicrocontrollerDocumentation(target, machineDocPath, co.controllerDocPath,
					createMachineDoc, updateBoardInterfaces, updateBoardPins); err != nil {
					fmt.Fprintln(os.Stderr, "  skipping: because error\t\t", err)
					co.errors = append(co.errors, err)
				}
			}
		}

		if createHardwareMatrix {
			fmt.Println(" Prepare hardware matrix...")
			matrixRow, err := hardwarematrix.ParseFeatureMap(co.controllerDocPath)
			if err != nil {
				fmt.Fprintln(os.Stderr, "  skipping: because error\t\t", err)
				co.errors = append(co.errors, err)
				continue
			}

			co.matrixRow = matrixRow
		}

		if co.docStatus == documentFound {
			co.docStatus = finishedWithoutTargetOk
			continue
		}
		co.docStatus = finishedOk
	}

	if createHardwareMatrix {
		fmt.Println("\nUpdate hardware matrix...")
		hwmatrixDocPath := filepath.Join("..", "content", "docs", "reference", "hardware-matrix.md")
		if err := updateHardwareMatrix(hwmatrixDocPath, collector); err != nil {
			fmt.Fprintln(os.Stderr, "  skipping: because error\t\t", err)
		}
	}

	// summary
	countErrors := printSummary(collector)

	if countErrors > 0 {
		os.Exit(1)
	}

	os.Exit(0)
}

func printSummary(collector map[string]*collection) int {
	// prepare the user output
	var countMissingDocsOk, countErrors, countFinishedOk int
	var targetsWithMissingDoc, targetsWithMissingDocButShouldHave []string
	var docWithMissingTarget []string
	for target, co := range collector {
		switch co.docStatus {
		case documentFound, finishedWithoutTargetOk:
			docWithMissingTarget = append(docWithMissingTarget, co.controllerDocPath)
		case missingDocOk:
			countMissingDocsOk++
		case missingDocButShouldHave:
			targetsWithMissingDocButShouldHave = append(targetsWithMissingDocButShouldHave, target)
		case missingDoc:
			targetsWithMissingDoc = append(targetsWithMissingDoc, target)
		}

		countErrors += len(co.errors)

		if co.docStatus == finishedOk {
			countFinishedOk++
		}
	}

	// show the user output
	fmt.Println("\n+++ results +++")
	fmt.Printf("%d\tdocuments found\n", len(collector))
	fmt.Printf("%d\ttargets known for missing document\n", countMissingDocsOk)
	fmt.Printf("%d\ttargets with missing documents\n", len(targetsWithMissingDoc))
	fmt.Printf("%d\ttargets with missing documents but should have one\n", len(targetsWithMissingDocButShouldHave))
	fmt.Printf("%d\ttargets with errors\n", countErrors)
	fmt.Printf("%d\ttargets finished successfully\n", countFinishedOk)
	fmt.Printf("%d\tpossible orphaned documents\n", len(docWithMissingTarget))
	if len(targetsWithMissingDoc) > 0 {
		fmt.Println("\n+++ a missing document is unexpected for this targets +++")
		sort.Strings(targetsWithMissingDoc)
		fmt.Println(strings.Join(targetsWithMissingDoc, ", "))
	}

	if len(docWithMissingTarget) > 0 {
		fmt.Println("\n+++ orphaned documents without target +++")
		sort.Strings(docWithMissingTarget)
		fmt.Println(strings.Join(docWithMissingTarget, ", "))
	}

	return countErrors
}

func collectFiles(targets []string) map[string]*collection {
	// collect values for processing and statistic, the key is the target
	collector := make(map[string]*collection)

	// find all relevant documents in the controller folder
	entries, err := os.ReadDir(controllerDocBaseDir)
	if err != nil {
		log.Fatal(err)
	}

	// collect all existing files
	for _, entry := range entries {
		name := entry.Name()
		extension := filepath.Ext(name)
		if entry.IsDir() || strings.HasPrefix(name, "_") || extension != ".md" {
			continue
		}

		target := strings.TrimSuffix(name, extension)
		co := collection{docStatus: documentFound, controllerDocPath: filepath.Join(controllerDocBaseDir, name)}
		collector[target] = &co
	}

	for _, target := range targets {
		co, ok := collector[target]
		if !ok {
			// document not exist
			co := collection{docStatus: missingDoc}
			collector[target] = &co

			if _, ok := targetsKnownToSkip[target]; ok {
				// this is not an error, so we skip it silently
				co.docStatus = missingDocOk
				continue
			}

			if _, ok := targetsKnownToSkipAndShouldBeDocumented[target]; ok {
				// this targets are known as missing, so we skip it here silently, but those are open TODO's
				co.docStatus = missingDocButShouldHave
				continue
			}

			fmt.Fprintln(os.Stderr, "Skipping: because missing microcontroller markdown document\t", target)
			continue
		}

		co.docStatus = canProcessed
	}

	return collector
}

func updateMicrocontrollerDocumentation(target, machineDocPath, controllerDocPath string,
	createMachineDoc, updateBoardInterfaces, updateBoardPins bool) error {
	ti, err := targetinfo.Create(target)
	if err != nil {
		return fmt.Errorf("target info can not be created: %v", err)
	}

	if createMachineDoc {
		err = machinepackagedoc.CreateNew(machineDocPath, target, ti.Pkg)
		if err != nil {
			return fmt.Errorf("error on write machine package doc: %v", err)
		}
	}

	errors := boarddoc.Update(controllerDocPath, ti, updateBoardInterfaces, updateBoardPins)
	if len(errors) > 0 {
		return fmt.Errorf("error on update board documentation: %v", errors)
	}

	return nil
}

func updateHardwareMatrix(hwmatrixDocPath string, collector map[string]*collection) error {
	var matrixRows []*hardwarematrix.RowInfo
	for _, co := range collector {
		if co.matrixRow == nil || co.matrixRow.Features == nil {
			continue
		}
		matrixRows = append(matrixRows, co.matrixRow)
	}

	if err := hardwarematrix.Create(hwmatrixDocPath, matrixRows); err != nil {
		return fmt.Errorf("hardware matrix can not be created: %v", err)
	}

	return nil
}
