package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/tools/go/packages"
)

// PackageDoc contains all the information necessary to render the documentation
// of a single package.
type PackageDoc struct {
	Target    string
	Types     map[string]*Type
	Funcs     map[string]*ast.FuncDecl
	Constants []*ast.GenDecl
	Variables []*ast.GenDecl
}

// Type an AST type plus all the functions that belong to it.
type Type struct {
	Type  *ast.TypeSpec
	Doc   *ast.CommentGroup
	Funcs map[string]*ast.FuncDecl
}

var markdownTemplateText = `
---
title: {{.Target}}
---

{{if .Constants}}
## Constants
{{range .Constants}}
{{format .}}
{{.Doc.Text}}
{{end}}
{{end}}

{{if .Variables}}
## Variables
{{range .Variables}}
{{format .}}
{{.Doc.Text}}
{{end}}
{{end}}

{{range $funcName, $func := .Funcs}}
### func {{$funcName}}

{{format $func}}
{{$func.Doc.Text}}
{{end}}

{{range $typeName, $type := .Types}}
## type {{$typeName}}

{{format $type.Type}}
{{$type.Doc.Text}}

{{range $funcName, $func := $type.Funcs }}
### func ({{formatReceiver $func.Recv}}) {{$funcName}}

{{format $func}}
{{$func.Doc.Text}}
{{end}}

{{end}}
`

func main() {
	// Get the target from the list of command line options.
	for _, target := range os.Args[1:] {
		path := filepath.Join("..", "content", "docs", "reference", "microcontrollers", "machine", target+".md")
		docPath := filepath.Join("..", "content", "docs", "reference", "microcontrollers", target+".md")
		if _, err := os.Stat(docPath); err != nil {
			fmt.Println("Skipping:                    ", target)
			continue
		}
		fmt.Println("Generating documentation for:", target)
		updateDocumentation(target, path, docPath)
	}
}

func updateDocumentation(target, path, docPath string) {
	// Get the important information from the compiler.
	cmd := exec.Command("tinygo", "info", target)
	outBytes, err := cmd.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not run tinygo:", err)
		os.Exit(1)
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
		fmt.Fprintln(os.Stderr, "could not find all needed properties (build tags, GOOS, GOARCH, GOROOT)")
		os.Exit(1)
	}

	// Get the list of files that would be compiled for this package.
	pkgs, err := packages.Load(&packages.Config{
		Mode:       packages.NeedName | packages.NeedSyntax | packages.NeedTypes | packages.NeedDeps,
		BuildFlags: []string{"-tags=" + strings.Join(buildTags, " ")},
		Env:        append(os.Environ(), "GOOS="+goos, "GOARCH="+goarch, "GOROOT="+goroot, "GO111MODULE=off"),
	}, "github.com/tinygo-org/tinygo/src/machine")
	if err != nil {
		log.Fatal(err)
	}

	// Do some sanity checking.
	if len(pkgs[0].Errors) != 0 {
		for _, err := range pkgs[0].Errors {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
	pkg := pkgs[0]

	err = writeMachinePackageDoc(path, target, pkg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	err = updateBoardDocumentation(docPath, pkg, buildTags)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func writeMachinePackageDoc(path, target string, pkg *packages.Package) error {
	doc := PackageDoc{
		Target: target,
		Types:  make(map[string]*Type),
		Funcs:  make(map[string]*ast.FuncDecl),
	}

	// Read everything except for functions.
	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			switch decl := decl.(type) {
			case *ast.GenDecl:
				if decl.Tok == token.CONST || decl.Tok == token.VAR {
					hasExportedNames := false
					for _, spec := range decl.Specs {
						for _, name := range spec.(*ast.ValueSpec).Names {
							if name.IsExported() {
								hasExportedNames = true
							}
						}
					}
					if hasExportedNames {
						switch decl.Tok {
						case token.CONST:
							doc.Constants = append(doc.Constants, decl)
						case token.VAR:
							doc.Variables = append(doc.Variables, decl)
						}
					}
					continue
				}
				for _, spec := range decl.Specs {
					switch spec := spec.(type) {
					case *ast.TypeSpec:
						if !spec.Name.IsExported() {
							continue
						}
						doc.Types[spec.Name.String()] = &Type{
							Type:  spec,
							Doc:   decl.Doc,
							Funcs: make(map[string]*ast.FuncDecl),
						}
					}
				}
			}
		}
	}

	// Read functions after types have been read, to attach functions to type
	// documentation.
	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			switch decl := decl.(type) {
			case *ast.FuncDecl:
				if !decl.Name.IsExported() {
					continue
				}
				if decl.Recv == nil {
					// TODO: check for return type
					doc.Funcs[decl.Name.Name] = decl
				} else {
					receiver := decl.Recv.List[0].Type
					var receiverName string
					switch receiver := receiver.(type) {
					case *ast.Ident:
						receiverName = receiver.Name
					case *ast.StarExpr:
						receiverName = receiver.X.(*ast.Ident).Name
					default:
						return fmt.Errorf("unknown receiver for %s", decl.Name.Name)
					}
					doc.Types[receiverName].Funcs[decl.Name.Name] = decl
				}
			}
		}
	}

	// Render the markdown of the documentation.
	tpl := template.Must(template.New("markdown").Funcs(map[string]interface{}{
		"format": func(node interface{}) string {
			w := &bytes.Buffer{}
			switch n := node.(type) {
			case *ast.FuncDecl:
				node = &ast.FuncDecl{
					Recv: n.Recv,
					Name: n.Name,
					Type: n.Type,
				}
			case *ast.TypeSpec:
				node = &ast.GenDecl{
					Tok:   token.TYPE,
					Specs: []ast.Spec{n},
				}
			case *ast.GenDecl:
				node = &ast.GenDecl{
					// TODO: filter unexported specs
					Tok:    n.Tok,
					Lparen: n.Lparen,
					Specs:  n.Specs,
					Rparen: n.Rparen,
				}
			}
			printer.Fprint(w, pkg.Fset, node)
			text := string(w.Bytes())
			return "```go\n" + text + "\n```\n"
		},
		"formatReceiver": func(receiver *ast.FieldList) string {
			// Special case for formatting the receiver of a method.
			w := &bytes.Buffer{}
			printer.Fprint(w, pkg.Fset, receiver.List[0].Type)
			return string(w.Bytes())
		},
	}).Parse(markdownTemplateText))

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer f.Close()

	return tpl.Execute(f, doc)
}

func updateBoardDocumentation(path string, pkg *packages.Package, buildTags []string) error {
	features := detectSupportedFeatures(pkg, buildTags)

	// Read the entire Markdown file in memory.
	docBuf, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not read Markdown file: %w", err)
	}
	doc := string(docBuf)

	// Find "Interfaces" section.
	start := strings.Index(doc, "## Interfaces\n")
	if start < 0 {
		return fmt.Errorf("could not find 'Interface' header")
	}
	endOffset := strings.Index(doc[start:], "\n## ")
	if endOffset < 0 {
		return fmt.Errorf("could not find end of 'Interface' header")
	}
	end := start + endOffset

	// Create new "Interfaces" section based on the previous section.
	interfacesText := "## Interfaces\n\n| Interface | Hardware Supported | TinyGo Support |\n| --------- | ------------- | ----- |\n"
	for _, line := range strings.Split(doc[start:end], "\n") {
		if !strings.HasPrefix(line, "| ") {
			continue
		}
		parts := strings.Split(line, "|")
		if len(parts) != 5 {
			return fmt.Errorf("expected 5 parts, got %d", len(parts))
		}
		feature := strings.TrimSpace(parts[1])
		if feature == "Interface" || feature[0] == '-' {
			// Part of the hearder.
			continue
		}
		hardwareSupport := strings.TrimSpace(parts[2])
		softwareSupport := strings.TrimSpace(parts[3])
		if hardwareSupport != "NO" && hardwareSupport != "?" {
			if supported, ok := features[feature]; ok {
				if supported {
					softwareSupport = "YES"
				} else {
					softwareSupport = "Not yet"
				}
			}
		}
		interfacesText += fmt.Sprintf("| %s      | %s | %s |\n", feature, hardwareSupport, softwareSupport)
	}

	// Replace the "Interfaces" section.
	doc = doc[:start] + interfacesText + doc[end:]

	// Write out the updated documentation.
	err = ioutil.WriteFile(path+".tmp", []byte(doc), 0o666)
	if err != nil {
		return fmt.Errorf("could not write updated Markdown file: %w", err)
	}
	err = os.Rename(path+".tmp", path)
	if err != nil {
		return fmt.Errorf("could not rename updated Markdown file: %w", err)
	}

	return nil
}

// detectSupportedFeatures check whether a given feature is supported by the
// chip/board based on the available types in the machine package and the build
// tags used for the compilation.
func detectSupportedFeatures(pkg *packages.Package, buildTags []string) map[string]bool {
	features := map[string]bool{
		"GPIO":      false,
		"UART":      false,
		"SPI":       false,
		"I2C":       false,
		"ADC":       false,
		"PWM":       false,
		"Bluetooth": false,
	}

	pinType := pkg.Types.Scope().Lookup("Pin").Type()
	if types.NewMethodSet(pinType).Lookup(pkg.Types, "Get") != nil {
		// Note: checking the 'Get' method because the 'Set' method is always
		// implemented (even if it's a no-op).
		features["GPIO"] = true
	}
	if pkg.Types.Scope().Lookup("UART") != nil {
		features["UART"] = true
	}
	if pkg.Types.Scope().Lookup("SPI") != nil {
		features["SPI"] = true
	}
	if pkg.Types.Scope().Lookup("I2C") != nil {
		features["I2C"] = true
	}
	adcType := pkg.Types.Scope().Lookup("ADC").Type()
	if types.NewMethodSet(adcType).Lookup(pkg.Types, "Configure") != nil {
		features["ADC"] = true
	}
	for _, tag := range buildTags {
		if tag == "nrf51" || tag == "nrf52" || tag == "nrf52840" || tag == "nrf52833" {
			features["Bluetooth"] = true
		}
	}

	// Detecting PWM support is a bit more tricky.
	// We basically iterate over all global variables and check whether they
	// have Configure method that takes a single PWMConfig struct.
	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			switch decl := decl.(type) {
			case *ast.GenDecl:
				if decl.Tok != token.VAR {
					continue
				}
				for _, spec := range decl.Specs {
					// Found a global variable.
					spec := spec.(*ast.ValueSpec)
					name := spec.Names[0].Name
					if !ast.IsExported(name) {
						continue
					}
					varType := pkg.Types.Scope().Lookup(name).Type()
					// Check for Configure method.
					configureMethod := types.NewMethodSet(varType).Lookup(pkg.Types, "Configure")
					if configureMethod == nil {
						continue
					}
					// Check whether it has just one parameter and this
					// parameter is of type PWMConfig.
					params := configureMethod.Type().(*types.Signature).Params()
					if params.Len() != 1 || params.At(0).Type() != pkg.Types.Scope().Lookup("PWMConfig").Type() {
						continue
					}
					// Yes, this looks like a PWM peripheral.
					features["PWM"] = true
				}
			}
		}
	}

	return features
}
