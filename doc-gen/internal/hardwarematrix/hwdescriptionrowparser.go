package hardwarematrix

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	extAst "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

const (
	supported          = "x"
	notSupported       = "-"
	partiallySupported = "p"
	notYetSupported    = "n"
	unknown            = "?"
	// there are some variants, so we compare in lower case
	boardDocSupported   = "yes"
	boardDocUnsupported = "no"
	boardDocNotYet      = "not yet"
)

type feature string

const (
	featGPIO   feature = "GPIO"
	featUART   feature = "UART"
	featSPI    feature = "SPI"
	featI2C    feature = "I2C"
	featADC    feature = "ADC"
	featPWM    feature = "PWM"
	featUSBDev feature = "USBDev"
	featBT     feature = "BT "
	featWiFi   feature = "WiFi"
	featIMU    feature = "IMU"
	featNP     feature = "NePx"
	featOther  feature = "other"
)

// if those feature was found in interfaces or pins table, it will be set to "supported"
// if found in text of section "Peripherals and Drivers", it will be set to "not yet supported", if not already
// set to "supported"
var knownFeatureCatchWords = map[feature][]string{
	featGPIO:   {"gpio"},
	featUART:   {"uart"},
	featSPI:    {"spi"},
	featI2C:    {"i2c"},
	featADC:    {"adc"},
	featPWM:    {"pwm"},
	featUSBDev: {"usbdev", "usbdevice"},
	featBT:     {"bluetooth", "bt"},
	featWiFi:   {"wifi", "wi-fi"},
	featIMU:    {"imu", "acceleration", "accelerometer", "gyroscope", "magnetometer", "rotation"},
	featNP:     {"neopixel", "WS2812"},
}

var supportedFeatureCatchWords = map[feature][]string{
	featBT:   {"NINA-W102", "nRF51422", "nRF51822", "nRF52832", "nrf52840"},
	featWiFi: {"NINA-W102"},
	featIMU:  {"LIS3DH", "LSM6DS3", "LSM6DS33", "LSM6DS3TR", "LSM6DOX", "LSM6DSOX", "LSM9DS1", "LSM303AGR"},
}

// because we currently not really check "other" feature, we mark those findings as "partially supported" in general
// TODO: maybe this can be parsed from "content/docs/reference/devices/_index.md"
var otherFeatureCatchWords = map[feature][]string{
	featOther: {"APDS9960", "button", "buttons", "humidity", "HTS221", "LoRa", "LoRaWAN", "LPS22HB", "microphone",
		"MicroSD", "MP34DT06JTR", "pressure", "proximity", "SD-Card", "sht3x", "temperature", "thermistor", "wan"},
}

type Link struct {
	Text        string
	Destination string
}

// key is the feature (e.g. the first cell content of a "Interfaces" table row), this will be the new table header
// the value is just the concatenation of all old cells
type features map[feature]string

type RowInfo struct {
	ControllerDocPath string // used as fallback for destination, if the link was not found
	FoundTitle        string // used as fallback for text, if the link was not found
	FoundLink         *Link
	Features          features
}

type hwDescriptionToRowInfoTransformer struct {
	source []byte
	// -- search variables --
	searchTitlePattern       *regexp.Regexp
	searchPeripheralsHeading string
	// interfaces table
	searchInterfaceTableHeaderColumnText string
	searchInterfaceTableColumnNumber     int
	rowInfo                              *RowInfo
	// pins table
	searchPinsTableHeaderColumnText string
	// -- AST walk storage --
	mcuAllContent           string
	peripheralsHeadingFound bool
	// table related
	interfacesTableFound              bool
	pinsTableFound                    bool
	columnNumberFound                 int
	currentTableColNumber             int
	currentTableNextFeature           string
	currentTableNextFeatureAllContent string
}

// ParseFeatureMap fill the feature map with this strategy:
//   - detect the first link (to the manufacturer board description), maybe detectable by the text inside the
//     [Board Name], which is the same like the "title"
//   - list in section "Peripherals" found: parse list - done by catchwords
//   - table in section "Interfaces" found: parse table
//   - "Pins" section found: skip (TODO)
//   - other section found: parse for some known tags (NINA-W102, nrf52840, ...)
func ParseFeatureMap(pathToControllerDescription string) (*RowInfo, error) {
	// Read the entire Markdown file in memory.
	// TODO: Parser().Parse() can also read from file
	mdBuf, err := os.ReadFile(pathToControllerDescription)
	if err != nil {
		return nil, fmt.Errorf("could not read Markdown file: %w", err)
	}

	rowInfo := RowInfo{
		ControllerDocPath: pathToControllerDescription,
		Features:          make(map[feature]string),
	}

	// Instantiate our transformer
	transformer := hwDescriptionToRowInfoTransformer{
		searchTitlePattern:                   regexp.MustCompile(`title: ["|'](.+)["|']`),
		searchInterfaceTableHeaderColumnText: "Interface",
		searchInterfaceTableColumnNumber:     1,
		searchPinsTableHeaderColumnText:      "Alternative names",
		rowInfo:                              &rowInfo,
		searchPeripheralsHeading:             "Peripherals and Drivers",
	}
	// multiple AST transformers are supported, will be run in order of priority
	prioritizedTransformer := util.Prioritized(&transformer, 0)

	// prepare the parser and parse to AST
	gm := goldmark.New(
		goldmark.WithExtensions(extension.Table),
		goldmark.WithParserOptions(parser.WithASTTransformers(prioritizedTransformer)),
	)

	_ = gm.Parser().Parse(text.NewReader(mdBuf))
	//docAst.Dump(mdBuf, 2) // for debugging purposes

	return &rowInfo, nil
}

// Transform implements goldmark.parser.ASTTransformer
func (t *hwDescriptionToRowInfoTransformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	// source contains the complete document, the AST objects only describes the positions in source
	t.source = reader.Source()

	err := ast.Walk(node, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		// Each node will be visited twice, once when it is first encountered (entering), and again
		// after all the node's children have been visited (if any).

		// try to get some features from text
		walkStatus, err := t.peripheralsTextToFeatures(node, entering)
		if err != nil || walkStatus != ast.WalkContinue {
			return walkStatus, err
		}

		// try to get the MCU title to find the right link text for the first column
		walkStatus, err = t.headingToMcuTitle(node, entering)
		if err != nil || walkStatus != ast.WalkContinue {
			return walkStatus, err
		}

		// try to get the destination for the link in the first column
		walkStatus, err = t.linkToLink(node, entering)
		if err != nil || walkStatus != ast.WalkContinue {
			return walkStatus, err
		}

		walkStatus, err = t.tablesToFeatures(node, entering)
		if err != nil || walkStatus != ast.WalkContinue {
			return walkStatus, err
		}

		return ast.WalkContinue, err
	})

	if err != nil {
		log.Fatal("Error encountered while transforming AST:", err)
	}
}

func (t *hwDescriptionToRowInfoTransformer) peripheralsTextToFeatures(node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	switch tnode := node.(type) {
	case *ast.Heading:
		content := string(tnode.Text(t.source))
		t.peripheralsHeadingFound = content == t.searchPeripheralsHeading
	case *ast.Text:
		content := string(tnode.Text(t.source))
		if t.peripheralsHeadingFound {
			t.rowInfo.setSupportedState(content, supportedFeatureCatchWords, supported)
			t.rowInfo.setSupportedState(content, knownFeatureCatchWords, notYetSupported)
			t.rowInfo.setSupportedState(content, otherFeatureCatchWords, partiallySupported)
		}
	}
	return ast.WalkContinue, nil
}

func (t *hwDescriptionToRowInfoTransformer) headingToMcuTitle(node ast.Node, entering bool) (ast.WalkStatus, error) {
	if t.rowInfo.FoundTitle != "" {
		return ast.WalkContinue, nil
	}

	switch tnode := node.(type) {
	case *ast.Text:
		if _, isHeading := tnode.Parent().(*ast.Heading); isHeading {
			if entering {
				//heading.Dump(t.source, 2) // for debugging purposes
				newContent := string(tnode.Text(t.source))
				if strings.HasPrefix(newContent, "weight:") {
					return ast.WalkContinue, nil
				}
				if tnode.SoftLineBreak() {
					t.mcuAllContent += newContent
				} else {
					t.mcuAllContent = newContent
				}
			} else {
				if tnode.SoftLineBreak() {
					// content is not complete yet, continue walk
					return ast.WalkContinue, nil
				}
				if mcu := t.searchTitlePattern.FindStringSubmatch(t.mcuAllContent); len(mcu) > 1 {
					t.rowInfo.FoundTitle = mcu[1]
					return ast.WalkContinue, nil
				}
			}
		}
	}
	return ast.WalkContinue, nil
}

func (t *hwDescriptionToRowInfoTransformer) linkToLink(node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering || t.rowInfo.FoundTitle == "" || t.rowInfo.FoundLink != nil {
		return ast.WalkContinue, nil
	}
	switch tnode := node.(type) {
	case *ast.Text:
		if link, isLink := tnode.Parent().(*ast.Link); isLink {
			//link.Dump(t.source, 2) // for debugging purposes
			content := string(tnode.Text(t.source))
			// content is usually shorter than the title, because often skip the vendor/manufacturer
			// nevertheless, for the hardware matrix we use the complete title rather than the shorten one
			if strings.Contains(t.rowInfo.FoundTitle, content) {
				t.rowInfo.FoundLink = &Link{Text: t.rowInfo.FoundTitle, Destination: string(link.Destination)}
			}
		}
	}
	return ast.WalkContinue, nil

}

func (t *hwDescriptionToRowInfoTransformer) tablesToFeatures(node ast.Node, entering bool) (ast.WalkStatus, error) {
	switch tnode := node.(type) {
	case *extAst.Table:
		t.interfacesTableFound = false
		t.pinsTableFound = false
		t.columnNumberFound = 0
	case *extAst.TableHeader:
		t.currentTableColNumber = 0
	case *extAst.TableRow:
		t.currentTableColNumber = 0
		if entering {
			t.currentTableNextFeature = ""
			t.currentTableNextFeatureAllContent = ""
		}
		if !entering && t.interfacesTableFound && t.currentTableNextFeature != "" {
			// on leaving the row, set the feature from found concatenated values, if not yet set to "supported"
			if featState, ok := t.rowInfo.Features[feature(t.currentTableNextFeature)]; !ok || featState != supported {
				t.rowInfo.Features[feature(t.currentTableNextFeature)] = convertBoardValue(t.currentTableNextFeatureAllContent)
			}
		}
	case *extAst.TableCell:
		if !entering {
			return ast.WalkSkipChildren, nil
		}
		t.currentTableColNumber++
		cellContent := string(tnode.Text(t.source))
		if _, isHeader := tnode.Parent().(*extAst.TableHeader); isHeader {
			if t.currentTableColNumber == t.searchInterfaceTableColumnNumber {
				if cellContent == t.searchInterfaceTableHeaderColumnText {
					t.interfacesTableFound = true
					t.columnNumberFound = t.currentTableColNumber
					return ast.WalkSkipChildren, nil
				}
			}
			if cellContent == t.searchPinsTableHeaderColumnText {
				t.pinsTableFound = true
				t.columnNumberFound = t.currentTableColNumber
			}
			// all other cells of header are not important yet
			return ast.WalkSkipChildren, nil
		}

		// a cell/column of row after the header reached
		if t.currentTableColNumber == t.columnNumberFound {
			// this column contains the feature name for interfaces table
			// or the list of features for pins table

			if t.interfacesTableFound {
				for feat, featCatchWords := range knownFeatureCatchWords {
					for _, featCatchWord := range featCatchWords {
						if strings.ToLower(featCatchWord) == strings.ToLower(cellContent) {
							t.currentTableNextFeature = string(feat)
							return ast.WalkSkipChildren, nil
						}
					}
				}
			}

			if t.pinsTableFound {
				t.rowInfo.setSupportedState(cellContent, knownFeatureCatchWords, supported)
			}

			return ast.WalkSkipChildren, nil
		}

		// cell/column reached after the column with possible feature
		if t.interfacesTableFound && t.currentTableNextFeature != "" {
			// add content of all further cells to the current
			t.currentTableNextFeatureAllContent += cellContent
		}
	}

	return ast.WalkContinue, nil
}

func (ri *RowInfo) setSupportedState(checkContent string, catchWords map[feature][]string, state string) {
	checkContent = strings.ToLower(checkContent)
	checkContent = strings.ReplaceAll(checkContent, ",", "")
	checkContent = strings.ReplaceAll(checkContent, "`", "")

	splitWords := strings.Split(checkContent, " ")
	for feat, featWords := range catchWords {
		if featState, ok := ri.Features[feat]; ok && featState == supported {
			// already set to supported
			continue
		}
		for _, word := range splitWords {
			for _, featWord := range featWords {
				featWord = strings.ToLower(featWord)
				if word == featWord {
					ri.Features[feat] = state
					continue
				}
			}
		}
	}
}

func convertBoardValue(val string) string {
	if strings.HasSuffix(strings.ToLower(val), boardDocNotYet) {
		return notYetSupported
	}

	if strings.HasSuffix(strings.ToLower(val), boardDocSupported) {
		return supported
	}

	if strings.HasPrefix(strings.ToLower(val), boardDocUnsupported) {
		return notSupported
	}

	return unknown
}
