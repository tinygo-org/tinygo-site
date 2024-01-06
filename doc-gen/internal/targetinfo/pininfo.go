package targetinfo

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"regexp"
	"strings"
)

// AnyPin holds information about an arbitrary pin
type AnyPin struct {
	HardwareName string
	OtherNames   []string
	Peripherals  map[string][]string
}

// PinsInfo holds information about all pins which are suitable for the documentation
type PinsInfo struct {
	HasPeripheral map[string]bool
	ExposedPins   []*AnyPin
}

// Match formats:
// - PA0 (AVR)
// - PA00 (Microchip SAM series)
// - P0_00 (nrf series),
// - GPIO0 (esp series, rp2040)
// - P00 (sifive)
var pinIsHardwareName = regexp.MustCompile("^(P[A-Z][0-9]+|P[0-1]_[0-9]{2}|GPIO[0-9]+|P[0-9]{2})$")

// CreatePinsInfo creates info about all pins contained in the target info package
func (ti *TargetInfo) CreatePinsInfo() (*PinsInfo, error) {
	// There is a pins section. Update it from the machine package.
	supportedPeripherals := map[string]bool{
		"PWM": true,
	}

	pkg := ti.Pkg

	// Find board pin names
	pinNames := make(map[string]uint64)
	hardwarePins := make(map[string]struct{})
	pinPeripherals := make(map[string][]string)
	var pinNamesSlice []string
	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			switch decl := decl.(type) {
			case *ast.GenDecl:
				if decl.Tok != token.CONST {
					continue
				}
				for _, spec := range decl.Specs {
					// Found a constant.
					spec := spec.(*ast.ValueSpec)
					name := spec.Names[0].Name
					if !ast.IsExported(name) {
						continue
					}
					obj := pkg.Types.Scope().Lookup(name)
					if obj.Type().String() != "machine.Pin" {
						continue
					}
					val, ok := constant.Uint64Val(obj.(*types.Const).Val())
					if !ok {
						return nil, fmt.Errorf("expected pin %s with value %s to be representable by uint64",
							name, obj.(*types.Const).Val())
					}
					pinNamesSlice = append(pinNamesSlice, name)
					pinNames[name] = val
					if isHardwarePin(name, spec.Values) {
						hardwarePins[name] = struct{}{}
						comment := strings.TrimSpace(spec.Comment.Text())
						if strings.HasPrefix(comment, "peripherals: ") {
							pinPeripherals[name] = strings.Split(strings.TrimPrefix(comment, "peripherals: "), ", ")
						}
					}
				}
			}
		}
	}
	if _, ok := pinNames["NoPin"]; !ok {
		return nil, fmt.Errorf("could not find NoPin constant")
	}

	pins := make(map[uint64]*AnyPin)
	pi := PinsInfo{
		HasPeripheral: make(map[string]bool),
	}

	for _, name := range pinNamesSlice {
		if name == "NoPin" || pinNames[name] == pinNames["NoPin"] {
			continue
		}
		num := pinNames[name]
		pin, ok := pins[num]
		if !ok {
			pin = &AnyPin{Peripherals: map[string][]string{}}
			pins[num] = pin
		}
		if _, ok := hardwarePins[name]; ok {
			// Hardware name
			if pin.HardwareName != "" {
				return nil, fmt.Errorf("duplicate hardware pin name: %s and %s", pin.HardwareName, name)
			}
			pin.HardwareName = name
			for _, peripheral := range pinPeripherals[name] {
				peripheralName := strings.SplitN(peripheral, " ", 2)[0]
				peripheralType := classifyPeripheral(pkg, peripheralName)
				if supportedPeripherals[peripheralType] {
					pin.Peripherals[peripheralType] = append(pin.Peripherals[peripheralType], peripheral)
					pi.HasPeripheral[peripheralType] = true
				}
			}
		} else {
			// Other name
			if len(pin.OtherNames) == 0 {
				pi.ExposedPins = append(pi.ExposedPins, pin)
			}
			pin.OtherNames = append(pin.OtherNames, name)
		}
	}

	return &pi, nil
}

// isHardwarePin return whether this pin looks like a hardware pin name. A hardware pin constant is a constant like PB02
// on the Arduino: a constant defined by the chip, and not the board (which would be pin 13, D13, or the LED pin).
func isHardwarePin(name string, values []ast.Expr) bool {
	if !pinIsHardwareName.MatchString(name) {
		return false
	}
	if len(values) == 0 {
		// Pin constant is probably part of a constant like this:
		//   const (
		//     GPIO0 Pin = iota
		//     GPIO1
		//     GPIO2
		//     // etc
		//   )
		return true
	}
	switch value := values[0].(type) {
	case *ast.BasicLit:
		// For example:
		//   const GPIO5 Pin = 5
		return true
	case *ast.BinaryExpr:
		// For example:
		//   const PB02 = portB + 2
		return true
	case *ast.Ident:
		if value.Name == "iota" {
			// Pins are initialized using the special identifier "iota".
			return true
		}
		// Doesn't look like a hardware pin, could be something like:
		//   const D13 = PB5
		return false
	default:
		return false
	}
}
