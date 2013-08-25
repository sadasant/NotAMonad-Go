package nonmonads

import (
	"fmt"
)

type Val interface{}
type Printer func(Val)

// Returns a printer that always prints the given value to the debug
// prompt.
//
// AWKWARD: The same value is always printed, regardless of what input
// value is given to the printer at the time.
func (p Printer) Wrap(v Val) Printer {
	return func(m Val) {
		fmt.Println(v.(string))
	}
}

// Returns a parser that uses the given parser but transforms its
// resulting value with the given transformation.
func (p Printer) Transform(t func(Val) Val) Printer {
	return func(m Val) {
		p(t(m))
	}
}

// Returns a printer that transforms values into a printer then prints
// that printer with the given printer printer.
func (p Printer) Flatten() Printer {
	return func(m Val) {
		p(p.Wrap(m))
	}
}
