package targetinfo

import (
	"go/ast"
	"go/token"
	"go/types"
)

// SupportedFeatures check whether a given feature is supported by the chip/board based on the available types
// in the machine package and the build tags used for the compilation.
func (ti *TargetInfo) SupportedFeatures() map[string]bool {
	features := map[string]bool{
		"GPIO":      false,
		"UART":      false,
		"SPI":       false,
		"I2C":       false,
		"ADC":       false,
		"PWM":       false,
		"Bluetooth": false,
		"USBDevice": false,
	}

	pkg := ti.Pkg

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
	for _, tag := range ti.BuildTags {
		if tag == "nrf51" || tag == "nrf52" || tag == "nrf52840" || tag == "nrf52833" {
			features["Bluetooth"] = true
		}
	}
	if pkg.Types.Scope().Lookup("USBDevice") != nil {
		features["USBDevice"] = true
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
					if classifyPeripheral(pkg, name) == "PWM" {
						features["PWM"] = true
					}
				}
			}
		}
	}

	return features
}
