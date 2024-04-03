package hardwarematrix

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

var markdownTemplate = `---
title: " Supported hardware feature matrix"
weight: 10
description: >
  Matrix between Microcontrollers and Features
---

Notes:
{{range $symbol, $description := .SupportedTypes}}
* '{{$symbol}}': {{$description}}{{end}}

see TinyGo documentation of the microcontroller for details to:

* WiFi
* BT (Bluetooth)
* IMU (Inertial measurement unit, e.g. acceleration, rotation, magnetometer)
* NePx (NeoPixel, WS2812)
* other (other peripheral or built-in devices, e.g. temperature, GSM)
{{range $_, $row := .Table}}
|{{range $_, $cell := $row}}{{$cell}}|{{end}}{{end}}
`

var supportedTypes = map[string]string{
	supported:          "Supported",
	notSupported:       "Not supported by Hardware",
	partiallySupported: "Partially supported by TinyGo",
	notYetSupported:    "Not yet supported by TinyGo",
	unknown:            "Supported status is currently unknown",
}

var headline = []string{
	" Microcontroller                             ", string(featGPIO), string(featUART), string(featSPI),
	string(featI2C), string(featADC), string(featPWM), string(featUSBDev), string(featBT), string(featWiFi),
	string(featIMU), string(featNP), string(featOther)}

// PackageDoc contains all the information necessary to render the matrix documentation table.
type PackageDoc struct {
	SupportedTypes map[string]string
	Table          [][]string
}

var linkTitlePattern = regexp.MustCompile(`\[(.+)\]`)

// Create creates a new document using the template.
func Create(pathToHardwareMatrix string, rowInfo []*RowInfo) error {
	doc := PackageDoc{
		SupportedTypes: supportedTypes,
	}

	doc.Table = append(doc.Table, headline, createSecondRow(headline))
	doc.Table = append(doc.Table, createSortedRows(headline, rowInfo, pathToHardwareMatrix)...)

	// Render the markdown of the documentation.
	tpl := template.Must(template.New("markdown").Parse(markdownTemplate))

	f, err := os.Create(pathToHardwareMatrix)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer f.Close()

	return tpl.Execute(f, doc)
}

func createSecondRow(headline []string) []string {
	row := make([]string, len(headline))
	for idx, cell := range headline {
		row[idx] = ":"
		for i := 0; i < len(cell)-2; i++ {
			row[idx] = row[idx] + "-"
		}
		row[idx] = row[idx] + ":"
	}

	return row
}

func createSortedRows(headline []string, rowInfos []*RowInfo, pathToHardwareMatrix string) [][]string {
	var rows [][]string
	firstColumn := make(map[string]struct{})
	for _, rowInfo := range rowInfos {
		row := make([]string, len(headline))
		for columnIdx, hlContent := range headline {
			content := notSupported
			length := len(hlContent)
			// strategy for the first row:
			// * if link exist, use it
			// * otherwise, create a link Text=title and Destination=link to md document
			// * if the result is not unique for the first column, add the title after link
			if columnIdx == 0 {
				if rowInfo.FoundLink == nil {
					if rowInfo.FoundTitle == "" {
						rowInfo.FoundTitle = "?????"
					}
					relPath, _ := filepath.Rel(pathToHardwareMatrix, rowInfo.ControllerDocPath)
					rowInfo.FoundLink = &Link{Text: rowInfo.FoundTitle, Destination: relPath}
				}
				if rowInfo.FoundLink != nil {
					content = fmt.Sprintf("[%s](%s)", rowInfo.FoundLink.Text, rowInfo.FoundLink.Destination)
				}
				// try to ensure unique content
				if _, ok := firstColumn[content]; ok {
					content = fmt.Sprintf("%s %s", content, rowInfo.FoundTitle)
				}
				firstColumn[content] = struct{}{}
				length = len(rowInfo.FoundLink.Destination) + length - 3 // means the visible length with collapsed link
			} else {
				if boardValue, ok := rowInfo.Features[feature(hlContent)]; ok {
					content = boardValue
				}
			}
			row[columnIdx] = fmt.Sprintf("%-*s", length, " "+content)
		}
		rows = append(rows, row)
	}

	// sort by first column
	sort.Slice(rows[:], func(i, j int) bool {
		for x := range rows[i] {
			a := strings.ToLower(rows[i][x])
			if link := linkTitlePattern.FindStringSubmatch(a); len(link) > 1 {
				a = link[1]
			}
			b := strings.ToLower(rows[j][x])
			if link := linkTitlePattern.FindStringSubmatch(b); len(link) > 1 {
				b = link[1]
			}
			if a == b {
				continue
			}
			return a < b
		}
		return false
	})

	return rows
}

// Update updates only what is changed, if a new controller is present, it will be added
// TODO: implement this update method - this can be done with "goldmark AST"
func Update(pathToHardwareMatrix string, rowInfo []*RowInfo) error {
	/*
		// TODO: goldmark.WithExtensions(extension.Table) is currently not supported by the renderer
		// prepare renderer and render to md
		mdRenderer := markdown.NewRenderer()
		buf := bytes.Buffer{} // implements io.Writer
		mdRenderer.Render(&buf, mdBufHwm, docAst)
	*/
	return nil
}
