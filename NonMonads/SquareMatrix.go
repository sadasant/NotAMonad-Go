package nonmonads

import (
	"errors"
)

type SquareMatrix [][]interface{}

func (s SquareMatrix) IsOk() error {
	if len(s[0]) != len(s[1]) {
		return errors.New("The SquareMatrix is not Well Formed")
	}
	return nil
}

// Returns a square matrix containing a single value: the given value.
func (s SquareMatrix) Wrap(v interface{}) (r SquareMatrix) {
	return SquareMatrix{[]interface{}{v}}
}

// Returns a square matrix created by running all of the items in the
// given square matrix through a transformation function.
func (s SquareMatrix) Transform(t func(interface{}) interface{}) SquareMatrix {
	r := make(SquareMatrix, len(s[0]))
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			r[i] = append(r[i], t(s[i][j]))
		}
	}
	return r
}

// Returns a square matrix containing the items inside the square
// matrices inside a square matrix.
// The items are ordered like you'd expect: by outer position then by
// inner position.
//
// BROKEN - When inner matrices are of different sizes, there may not be
// a square number of items.
func (s SquareMatrix) Flatten() SquareMatrix {
	switch len(s) {
	case 0:
		return s
	case 1:
		if s0, ok := s[0][0].(SquareMatrix); ok {
			return s0
		}
		return s
	}
	r := SquareMatrix{}
	for si := 0; si < len(s); si++ {
		for sj := 0; sj < len(s[si]); sj++ {
			if s0, ok := s[si][sj].(SquareMatrix); ok {
				innerSpan := len(s0)
				for s0i := 0; s0i < len(s0); s0i++ {
					for s0j := 0; s0j < len(s0[s0i]); s0j++ {
						i := si*innerSpan + s0i
						if len(r) <= i {
							r = append(r, []interface{}{})
						}
						j := sj*innerSpan + s0j
						if len(r[i]) <= j {
							r[i] = append(r[i], []interface{}{})
						}
						r[i][j] = s0[s0i][s0j]
					}
				}
			} else {
				r[si] = append(r[si], s[si][sj])
			}
		}
	}
	return r
}
