package machinepackagedoc

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	"text/template"

	"golang.org/x/tools/go/packages"
)

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

// Type an AST type plus all the functions that belong to it.
type Type struct {
	Type  *ast.TypeSpec
	Doc   *ast.CommentGroup
	Funcs map[string]*ast.FuncDecl
}

// PackageDoc contains all the information necessary to render the documentation
// of a single package.
type PackageDoc struct {
	Target    string
	Types     map[string]*Type
	Funcs     map[string]*ast.FuncDecl
	Constants []*ast.GenDecl
	Variables []*ast.GenDecl
}

// CreateNew writes the machine package documentation new from scratch
func CreateNew(path, target string, pkg *packages.Package) error {
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
					if ast.IsExported(receiverName) {
						doc.Types[receiverName].Funcs[decl.Name.Name] = decl
					}
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
