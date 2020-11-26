package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
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
		fmt.Println("Generating documentation for:", target)
		path := filepath.Join("..", "content", "microcontrollers", "machine", target+".md")
		writeTargetDoc(target, path)
	}
}

func writeTargetDoc(target, path string) {
	// Get the important information from the compiler.
	cmd := exec.Command("tinygo", "info", target)
	outBytes, err := cmd.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not run tinygo:", err)
		os.Exit(1)
	}
	var buildTags []string
	var goos, goarch string
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
		}
	}
	if len(buildTags) == 0 || goos == "" || goarch == "" {
		fmt.Fprintln(os.Stderr, "could not find all needed properties (build tags, GOOS, GOARCH)")
		os.Exit(1)
	}

	f, err := os.Create(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not open file:", err)
		os.Exit(1)
	}
	defer f.Close()
	writeDoc(f, target, buildTags, goos, goarch)
}

func writeDoc(out io.Writer, target string, buildTags []string, goos, goarch string) {
	// Get the list of files that would be compiled for this package.
	pkgs, err := packages.Load(&packages.Config{
		Mode:       packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles,
		BuildFlags: []string{"-tags=" + strings.Join(buildTags, " ")},
		Env:        append(os.Environ(), "GOOS="+goos, "GOARCH="+goarch, "GO111MODULE=off"),
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
	if len(pkgs[0].CompiledGoFiles) == 0 {
		fmt.Fprintln(os.Stderr, "no compiled Go files found for target:", target)
		os.Exit(1)
	}

	// Parse all to-be-compiled files.
	syntax := make([]*ast.File, len(pkgs[0].CompiledGoFiles))
	fset := token.NewFileSet()
	for i, fileName := range pkgs[0].CompiledGoFiles {
		file, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
		if err != nil {
			log.Fatal("failed to parse Go source file:", err)
		}
		syntax[i] = file
	}

	doc := PackageDoc{
		Target: target,
		Types:  make(map[string]*Type),
		Funcs:  make(map[string]*ast.FuncDecl),
	}

	// Read everything except for functions.
	for _, file := range syntax {
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
	for _, file := range syntax {
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
						log.Fatal("unknown receiver for:", decl.Name.Name)
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
			printer.Fprint(w, fset, node)
			text := string(w.Bytes())
			return "```go\n" + text + "\n```\n"
		},
		"formatReceiver": func(receiver *ast.FieldList) string {
			// Special case for formatting the receiver of a method.
			w := &bytes.Buffer{}
			printer.Fprint(w, fset, receiver.List[0].Type)
			return string(w.Bytes())
		},
	}).Parse(markdownTemplateText))
	err = tpl.Execute(out, doc)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
}
