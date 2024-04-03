package boarddoc

import (
	"fmt"
	"os"
	"strings"

	"github.com/tinygo-org/tinygo-site/doc-gen/internal/targetinfo"
)

func Update(path string, ti *targetinfo.TargetInfo, updateInterfaces, updatePins bool) []error {
	// Read the entire Markdown file in memory.
	docBuf, err := os.ReadFile(path)
	if err != nil {
		return []error{fmt.Errorf("could not read Markdown file: %w", err)}
	}
	doc := string(docBuf)

	var interfacesUpdated, pinsUpdated bool
	var errors []error
	if updateInterfaces {
		features := ti.SupportedFeatures()
		doc, interfacesUpdated, err = updateInterfacesSection(doc, features)
		if err != nil {
			fmt.Fprintln(os.Stderr, "interfaces section could not be updated", err)
			errors = append(errors, err)
		}
	}

	if updatePins {
		doc, pinsUpdated, err = updatePinsSection(doc, ti)
		if err != nil {
			fmt.Fprintln(os.Stderr, "pins section could not be updated", err)
			errors = append(errors, err)
		}
	}

	if !interfacesUpdated && !pinsUpdated {
		return errors
	}

	// Write out the updated documentation.
	err = os.WriteFile(path+".tmp", []byte(doc), 0o666)
	if err != nil {
		errors = append(errors, fmt.Errorf("could not write updated Markdown file: %w", err))
		return errors
	}
	err = os.Rename(path+".tmp", path)
	if err != nil {
		errors = append(errors, fmt.Errorf("could not rename updated Markdown file: %w", err))
		return errors
	}

	return nil
}

func updateInterfacesSection(doc string, features map[string]bool) (string, bool, error) {
	start, end, err := findSection("Interfaces", doc)
	if err != nil {
		// mandatory so return the error
		return doc, false, err
	}

	// Create new "Interfaces" section based on the previous one.
	interfacesText := "## Interfaces\n\n" +
		"| Interface | Hardware Supported | TinyGo Support |\n" +
		"| --------- | ------------- | ----- |\n"

	for _, line := range strings.Split(doc[start:end], "\n") {
		if !strings.HasPrefix(line, "| ") {
			continue
		}
		parts := strings.Split(line, "|")
		if len(parts) != 5 {
			return doc, false, fmt.Errorf("expected 5 parts, got %d", len(parts))
		}
		feature := strings.TrimSpace(parts[1])
		if feature == "Interface" || feature[0] == '-' {
			// Part of the header.
			continue
		}

		// If the feature is really supported by hardware, we set the software support
		// for all of known features. Unknown features keep on the manual set state.
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
		interfacesText += fmt.Sprintf("| %-9s | %-3s | %-3s |\n", feature, hardwareSupport, softwareSupport)
	}

	// Replace the "Interfaces" section.
	return doc[:start] + interfacesText + doc[end:], true, nil
}

func updatePinsSection(doc string, ti *targetinfo.TargetInfo) (string, bool, error) {
	start, end, err := findSection("Pins", doc)
	if err != nil {
		// optional, so no error
		fmt.Println("pins section not found - no change", err)
		return doc, false, nil
	}

	pinsInfo, err := ti.CreatePinsInfo()
	if err != nil {
		return doc, false, err
	}

	pinsText := "## Pins\n\n" +
		"| Pin               | Hardware pin | Alternative names |"

	if pinsInfo.HasPeripheral["PWM"] {
		pinsText += " PWM                  |"
	}
	pinsText += "\n| ----------------- | ------------ | ----------------- |"
	for range pinsInfo.HasPeripheral {
		pinsText += " -------------------- |"
	}
	pinsText += "\n"
	for _, pin := range pinsInfo.ExposedPins {
		if pin.HardwareName == "" {
			return doc, false, fmt.Errorf("could not find hardware pin name for %s", pin.OtherNames[0])
		}
		alternativeNames := make([]string, 0, len(pin.OtherNames)-1)
		for _, name := range pin.OtherNames[1:] {
			alternativeNames = append(alternativeNames, "`"+name+"`")
		}
		pinsText += fmt.Sprintf("| %-17s | %-12s | %-17s |", "`"+pin.OtherNames[0]+"`", "`"+pin.HardwareName+"`",
			strings.Join(alternativeNames, ", "))
		for _, peripheralType := range []string{"PWM"} {
			if !pinsInfo.HasPeripheral[peripheralType] {
				continue
			}
			var peripherals []string
			for _, peripheral := range pin.Peripherals[peripheralType] {
				parts := strings.SplitN(peripheral, " ", 2)
				s := "`" + parts[0] + "` (" + parts[1] + ")"
				s = strings.ReplaceAll(s, " ", "\u00a0") // use non-breaking space
				peripherals = append(peripherals, s)
			}
			pinsText += fmt.Sprintf(" %-20s |", strings.Join(peripherals, ", "))
		}
		pinsText += "\n"
	}

	// Replace the "Pins" section.
	return doc[:start] + pinsText + doc[end:], true, nil
}

func findSection(name, doc string) (start, end int, err error) {
	start = strings.Index(doc, fmt.Sprintf("## %s\n", name))
	if start < 0 {
		return 0, 0, fmt.Errorf("could not find '%s' header", name)
	}
	endOffset := strings.Index(doc[start:], "\n## ")
	if endOffset < 0 {
		return 0, 0, fmt.Errorf("could not find end of '%s' header", name)
	}
	end = start + endOffset
	return
}
