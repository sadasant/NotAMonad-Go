package nonmonads

import (
	"errors"
)

type Square [][]Val
type SquareMatrix struct {
	Items Square
}

func (s *SquareMatrix) From(square Square) (SquareMatrix, error) {
	s.Items = square
	err := s.isWellFormed()
	return *s, err
}

func (s SquareMatrix) isWellFormed() error {
	if len(s.Items) > 0 && len(s.Items[0]) != len(s.Items[1]) {
		return errors.New("The SquareMatrix is not Well Formed")
	}
	return nil
}

// Returns a square matrix containing a single value: the given value.
func (s SquareMatrix) Wrap(v Val) (r SquareMatrix) {
	r.Items = Square{[]Val{v}}
	return
}

// Returns a square matrix created by running all of the items in the
// given square matrix through a transformation function.
func (s SquareMatrix) Transform(t func(Val) Val) (SquareMatrix, error) {
	items := make(Square, len(s.Items[0]))
	for i := 0; i < len(s.Items); i++ {
		for j := 0; j < len(s.Items[i]); j++ {
			items[i] = append(items[i], t(s.Items[i][j]))
		}
	}
	return new(SquareMatrix).From(items)
}

// Returns a square matrix containing the items inside the square
// matrices inside a square matrix.
// The items are ordered like you'd expect: by outer position then by
// inner position.
//
// BROKEN - When inner matrices are of different sizes, there may not be
// a square number of items.
func (s SquareMatrix) Flatten() (SquareMatrix, error) {
	switch len(s.Items) {
	case 0:
		return s, nil
	case 1:
		if s0, ok := s.Items[0][0].(SquareMatrix); ok {
			return s0, nil
		}
		return s, nil
	}
	r := Square{}
	for si := 0; si < len(s.Items); si++ {
		for sj := 0; sj < len(s.Items[si]); sj++ {
			if s0, ok := s.Items[si][sj].(SquareMatrix); ok {
				innerSpan := len(s0.Items)
				for s0i := 0; s0i < len(s0.Items); s0i++ {
					for s0j := 0; s0j < len(s0.Items[s0i]); s0j++ {
						i := si*innerSpan + s0i
						if len(r) <= i {
							r = append(r, []Val{})
						}
						j := sj*innerSpan + s0j
						if len(r[i]) <= j {
							r[i] = append(r[i], []Val{})
						}
						r[i][j] = s0.Items[s0i][s0j]
					}
				}
			} else {
				r[si] = append(r[si], s.Items[si][sj])
			}
		}
	}
	return new(SquareMatrix).From(r)
}
