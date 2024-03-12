package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

const (
	testGOOS   = "linux"
	testGOARCH = "amd64"
)

// The environment to pass to `go` commands when they are invoked.
var commandEnv = []string{"GOPATH=/does-not-exist", "GO111MODULE=off", "GOOS=" + testGOOS, "GOARCH=" + testGOARCH}

var markdownTemplate = template.Must(template.New("markdown").Parse(`
---
title: Packages supported by TinyGo
---

The following table shows all Go standard library packages and whether they can be imported by TinyGo. If they can't, you can click the 'no' link to jump to the explanation why the package cannot be compiled.

Note that the fact they can be imported, does not mean that all functions and types in the program can be used. For example, sometimes using some functions or types of the package will still trigger compiler errors.

Test results are for {{.goos}}/{{.goarch}}.

Package | Importable | Passes tests
--- | --- | --- |{{ range .pkgs}}
{{.Path}} | {{if .CanBeCompiled}} <span style="color: green">✔</span> yes {{else}} [<span style="color: red">✗</span> no](#{{.Link}}) {{end}} | {{if .PassesTests}} <span style="color: green">✔</span> yes {{else if .CanBeCompiled}} [<span style="color: red">✗</span> no](#{{.Link}}) {{else}} <span style="color: gray">✗</span> no {{end}} | {{ end }}

{{range .pkgs}}
{{if not .PassesTests }}
## {{.Path}}

{{if .Output}}
{{if .CanBeCompiled}}
The compiler gave the following error when running the tests for this package:
{{else}}
The compiler gave the following error when this package was imported:
{{end}}

{{.IndentedOutput}}
{{else if .Imports}}
This package cannot be imported because the following dependencies cannot be compiled:
{{range .Imports}}{{if not .CanBeCompiled}}
  * [{{.Path}}](#{{.Link}}){{end}}{{end}}
{{else}}Could not compile for an unknown reason.
{{end}}

{{end}}
{{end}}
`))

type Package struct {
	Path          string
	Imports       []*Package
	Output        string
	CanBeCompiled bool
	PassesTests   bool
	Compiling     bool
}

func (p *Package) AllImportsCanCompile() bool {
	for _, pkg := range p.Imports {
		if !pkg.CanBeCompiled {
			return false
		}
	}
	return true
}

// Link returns the link name to a header with the path as title. Hugo/Docsy removes
// slashes so we emulate that behavior.
func (p *Package) Link() string {
	return strings.Replace(p.Path, "/", "", -1)
}

// IndentedOutput returns the Output variable but indented by 4 characters, so
// that it is valid Markdown raw output.
func (p *Package) IndentedOutput() string {
	// Remove last newline.
	output := strings.TrimSpace(p.Output)
	// Add indentation for all lines (except the last of course).
	output = "    " + strings.ReplaceAll(output, "\n", "\n    ")
	// Add last newline again.
	output += "\n"

	return output
}

// testResult is returned by runTest. It's necessary to communicate the test
// result back over a channel from a separate test runner goroutine.
type testResult struct {
	pkg            *Package
	compiles       bool
	compilerOutput string
	passesTest     bool
	testOutput     string
}

func checkPackages(goroot string) error {
	fmt.Fprintln(os.Stderr, "gathering a list of packages...")

	baseDir := filepath.Join(goroot, "src")

	// Get a list of all standard library packages.
	pkgsList := new(bytes.Buffer)
	cmd := exec.Command("go", "list", "std")
	cmd.Env = append(os.Environ(), commandEnv...)
	cmd.Dir = baseDir
	cmd.Stdout = pkgsList
	if err := cmd.Run(); err != nil {
		return err
	}

	// Get a list of all packages to inspect with their direct dependencies.
	pkgs := []*Package{}
	pkgMap := map[string]*Package{}
	for _, path := range strings.Split(string(pkgsList.Bytes()), "\n") {
		if !shouldTestPackage(path) {
			// Don't count this package.
			continue
		}

		// Valid package, add to the list of packages.
		pkg := &Package{
			Path: path,
		}
		pkgs = append(pkgs, pkg)
		pkgMap[path] = pkg
	}

	fmt.Fprintln(os.Stderr, "finding dependencies of each package...")

	// Get the dependencies of each package.
	for _, pkg := range pkgs {
		cmd := exec.Command("go", "list", "-f", `{{ join .Imports "\n" }}`, pkg.Path)
		buf := new(bytes.Buffer)
		cmd.Dir = baseDir
		cmd.Env = append(os.Environ(), commandEnv...)
		cmd.Stdout = buf
		err := cmd.Run()
		if err != nil {
			return err
		}
		for _, importPath := range strings.Split(strings.TrimSpace(string(buf.Bytes())), "\n") {
			if !shouldTestPackage(importPath) {
				continue
			}
			if _, ok := pkgMap[importPath]; !ok {
				return fmt.Errorf("unknown import %#v in package %s", importPath, pkg.Path)
			}
			pkg.Imports = append(pkg.Imports, pkgMap[importPath])
		}
	}

	testInput := make(chan *Package, 1)
	testOutput := make(chan testResult)
	for i := runtime.NumCPU(); i != 0; i-- {
		go func() {
			for pkg := range testInput {
				testOutput <- pkg.runTest()
			}
		}()
	}

	// Keep track how many tests are in flight, reaching zero means all tests
	// have been run.
	var testsInFlights int

	// Try to compile each package of which all imports can be compiled.
	for {
		// Fill pipeline of packages, as far as possible (possibly zero at the
		// end).
		for {
			inputPkg := nextPackageToTest(pkgs)
			if inputPkg == nil {
				// All packages have been tested.
				break
			}

			success := false
			select {
			case testInput <- inputPkg:
				inputPkg.Compiling = true
				success = true
				testsInFlights++
			default:
			}
			if !success {
				break
			}
		}

		if testsInFlights == 0 {
			break
		}

		// Receive one result and process it.
		result := <-testOutput
		result.pkg.Compiling = false
		testsInFlights--
		if !result.compiles {
			// There was an error importing this package.
			fmt.Fprintf(os.Stderr, "%-30s does not compile\n", result.pkg.Path)
			msg := result.compilerOutput
			if len(msg) == 0 {
				msg = "[...no output...]" // should not happen
			}
			lines := strings.Split(msg, "\n")
			if len(lines) > 15 {
				msg = strings.Join(lines[:15], "\n") + "\n[...more lines following...]"
			}
			result.pkg.Output = msg
		} else if !result.passesTest {
			// There was an error running the tests for this package.
			fmt.Fprintf(os.Stderr, "%-30s fails tests\n", result.pkg.Path)
			msg := result.testOutput
			if len(msg) == 0 {
				msg = "[...no test output...]" // should definitely not happen
			}
			lines := strings.Split(msg, "\n")
			if len(lines) > 15 {
				msg = strings.Join(lines[:15], "\n") + "\n[...more lines following...]"
			}
			result.pkg.Output = msg
			result.pkg.CanBeCompiled = true
		} else {
			// This package passes all tests!
			fmt.Fprintf(os.Stderr, "%-30s ok\n", result.pkg.Path)
			result.pkg.CanBeCompiled = true
			result.pkg.PassesTests = true
		}
	}

	// Print the output to stdout.
	return markdownTemplate.Execute(os.Stdout, map[string]interface{}{
		"pkgs":   pkgMap,
		"goos":   testGOOS,
		"goarch": testGOARCH,
	})
}

func (pkg *Package) runTest() (result testResult) {
	result.pkg = pkg

	// Prepare test files.
	dir, err := os.MkdirTemp("", "tinygo-test-*")
	if err != nil {
		panic("could not create temporary directory: " + err.Error())
	}
	defer os.RemoveAll(dir)
	temporaryGoFile := filepath.Join(dir, "main.go")
	os.WriteFile(temporaryGoFile, []byte(fmt.Sprintf("package main\nimport _ \"%s\"\nfunc main(){}\n", pkg.Path)), 0600)
	temporaryExecutableFile := filepath.Join(dir, "main")

	// Run the compile test.
	cmd := exec.Command("tinygo", "build", "-o", temporaryExecutableFile, temporaryGoFile)
	cmd.Env = append(cmd.Environ(), "GOOS="+testGOOS, "GOARCH="+testGOARCH)
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = buf
	result.compiles = cmd.Run() == nil
	result.compilerOutput = string(buf.Bytes())

	// Run the actual test.
	if result.compiles {
		cmd := exec.Command("tinygo", "test", pkg.Path)
		cmd.Env = append(cmd.Environ(), "GOOS="+testGOOS, "GOARCH="+testGOARCH)
		buf := new(bytes.Buffer)
		cmd.Stdout = buf
		cmd.Stderr = buf
		result.passesTest = cmd.Run() == nil
		result.testOutput = string(buf.Bytes())
	}

	return
}

// nextPackageToTest returns the next package to test, or nil if there is no
// available package to test.
// While there are tests in flight, this function might return nil even though
// there are packages available for testing.
func nextPackageToTest(pkgs []*Package) *Package {
	for _, pkg := range pkgs {
		if pkg.CanBeCompiled || pkg.Compiling || pkg.Output != "" || !pkg.AllImportsCanCompile() {
			// Already tested or one of the dependencies failed.
			continue
		}
		return pkg
	}
	return nil
}

func shouldTestPackage(path string) bool {
	if path == "" {
		return false
	}

	if path == "C" {
		// Skip CGo pseudo-package.
		return false
	}

	if path == "runtime" || strings.HasPrefix(path, "runtime/") {
		// Compiler-specific package.
		return false
	}

	if strings.HasPrefix(path, "vendor/") {
		// vendored package.
		return false
	}

	if strings.HasPrefix(path, "cmd/") {
		// Contains commands and not packages, and should not be importable
		// anyway.
		return false
	}

	for _, component := range strings.Split(path, "/") {
		if component == "internal" {
			// Internal package, not importable.
			return false
		}
	}

	// This package should be tested.
	return true
}

func main() {
	err := checkPackages(runtime.GOROOT())
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not check packages:", err)
		os.Exit(1)
	}
}
