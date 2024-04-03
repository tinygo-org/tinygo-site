package targetinfo

import (
	"fmt"
	"go/types"
	"log"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/tools/go/packages"
)

// TargetInfo contains info attributes of the target
type TargetInfo struct {
	BuildTags []string
	Pkg       *packages.Package
}

// Create retrieves some information about the given target
func Create(target string) (*TargetInfo, error) {
	// Get the important information from the compiler.
	cmd := exec.Command("tinygo", "info", target)
	outBytes, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("could not run tinygo: %v", err)
	}

	var buildTags []string
	var goos, goarch, goroot string
	for _, line := range strings.Split(string(outBytes), "\n") {
		idx := strings.IndexByte(line, ':')
		if idx < 0 {
			continue
		}
		key := strings.TrimSpace(line[:idx])
		value := strings.TrimSpace(line[idx+1:])
		switch key {
		case "build tags":
			buildTags = strings.Fields(value)
		case "GOOS":
			goos = value
		case "GOARCH":
			goarch = value
		case "cached GOROOT":
			goroot = value
		}
	}
	if len(buildTags) == 0 || goos == "" || goarch == "" || goroot == "" {
		return nil, fmt.Errorf("could not find all needed properties (build tags, GOOS, GOARCH, GOROOT)")
	}

	// Get the list of files that would be compiled for this package.
	pkgs, err := packages.Load(&packages.Config{
		Mode:       packages.NeedName | packages.NeedSyntax | packages.NeedTypes | packages.NeedDeps,
		BuildFlags: []string{"-tags=" + strings.Join(buildTags, " ")},
		Env:        append(os.Environ(), "GOOS="+goos, "GOARCH="+goarch, "GOROOT="+goroot, "GO111MODULE=off"),
	}, "machine")
	if err != nil {
		log.Fatal(err)
	}

	// Do some sanity checking.
	if len(pkgs[0].Errors) != 0 {
		fmt.Fprintf(os.Stderr, " sanity check failed with %d error(s)\n", len(pkgs[0].Errors))
		for _, err := range pkgs[0].Errors {
			fmt.Fprintf(os.Stderr, " * %v\n", err)
		}
		return nil, fmt.Errorf("sanity check failed")
	}

	ti := TargetInfo{
		Pkg:       pkgs[0],
		BuildTags: buildTags,
	}

	return &ti, nil
}

// classifyPeripheral returns the peripheral type for a given peripheral name, by looking at the type of the global.
func classifyPeripheral(pkg *packages.Package, name string) string {
	global := pkg.Types.Scope().Lookup(name)
	if global == nil {
		// The peripheral is not available.
		// This sometimes happen when the pins are defined for a chip family but
		// some chips don't have all the peripherals.
		return ""
	}

	// Check for Configure method. Most peripherals have one.
	configureMethod := types.NewMethodSet(global.Type()).Lookup(pkg.Types, "Configure")
	if configureMethod == nil {
		return ""
	}

	// Check whether it has just one parameter and this parameter is of type
	// PWMConfig.
	params := configureMethod.Type().(*types.Signature).Params()
	if params.Len() == 1 && params.At(0).Type() == pkg.Types.Scope().Lookup("PWMConfig").Type() {
		return "PWM"
	}
	// Some other kind of peripheral.
	return ""
}
