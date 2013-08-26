package nonmonads

import (
	"fmt"
)

type Printer func(interface{})

// Returns a printer that always prints the given value to the debug
// prompt.
//
// AWKWARD: The same value is always printed, regardless of what input
// value is given to the printer at the time.
func (p Printer) Wrap(v interface{}) Printer {
	return func(m interface{}) {
		fmt.Println(v.(string))
	}
}

// Returns a parser that uses the given parser but transforms its
// resulting value with the given transformation.
func (p Printer) Transform(t func(interface{}) interface{}) Printer {
	return func(m interface{}) {
		p(t(m))
	}
}

// Returns a printer that transforms values into a printer then prints
// that printer with the given printer printer.
func (p Printer) Flatten() Printer {
	return func(m interface{}) {
		p(p.Wrap(m))
	}
}
