package nonmonads

import (
	"errors"
	"math/cmplx"
)

type Superposition map[string]interface{}

type QS func() Superposition

// well-formed superpositions must add up to 100%:
func (p Superposition) IsOk() error {
	var sum float64
	for _, v := range p {
		if _v, ok := v.([]interface{}); ok {
			sum += cmplx.Abs(_v[1].(complex128))
		} else {
			sum += cmplx.Abs(v.(complex128))
		}
	}
	if sum != 1 {
		return errors.New("The QS is not Well Formed")
	}
	return nil
}

// Returns a superposition where the given value is the single state and
// has an amplitude of 1.
func (q QS) Wrap(k string) QS {
	return func() Superposition {
		return Superposition{k: complex(1.0, 0)}
	}
}

// Returns a superposition over the result of drawing an input from the
// given superposition, then running that input through the given
// transformation.
//
// BROKEN - When distinct inputs are merged, they interfere.
// The interference breaks the squared magnitude constraint.
func (q QS) Transform(t func(string) string) QS {
	qs := q()
	rs := Superposition{}
	for k, v := range qs {
		trans := t(k)
		if _, ok := rs[trans]; ok {
			rs[trans] = rs[trans].(complex128) + v.(complex128)
		} else {
			if _v, ok := v.([]interface{}); ok {
				rs[trans] = _v[1]
			} else {
				rs[trans] = v
			}
		}
	}
	return func() Superposition { return rs }
}

// Returns a superposition over the result of drawing an intermediate
// superposition from the given superposition of superpositions,
// then drawing an item from that intermediate superposition.
//
// BROKEN - When an item appears in multiple intermediate superpositions,
// the squared magnitude constraint may be violated.
//
// ISSUE:
// If we don't use our custom mult function, it breaks
// In go1.1.2 linux/arm it breaks with: reg R13 left allocated
// The issue was reported here: https://code.google.com/p/go/issues/detail?can=2&start=0&num=100&q=&colspec=ID%20Status%20Stars%20Priority%20Owner%20Reporter%20Summary&groupby=&sort=&id=6247
func (q QS) Flatten() QS {
	qs := q()
	rs := Superposition{}
	for _, l := range qs {
		if _l, ok := l.([]interface{}); ok {
			for k, v := range _l[0].(QS)() {
				if _, ok := rs[k]; ok {
					rs[k] = rs[k].(complex128) + mult(v.(complex128), _l[1].(complex128))
				} else {
					if _v, ok := v.([]interface{}); ok {
						rs[k] = mult(_v[1].(complex128), _l[1].(complex128))
					} else {
						rs[k] = mult(v.(complex128), _l[1].(complex128))
					}
				}
			}
		} else {
			return q
		}
	}
	return func() Superposition { return rs }
}

// As appears here: http://www.clarku.edu/~djoyce/complex/mult.html
func mult(a, b complex128) complex128 {
	ra := real(a)
	rb := real(b)
	ia := imag(a)
	ib := imag(b)
	r := (ra * rb) - (ia * ib)
	i := (ra * ib) + (rb * ia)
	// Becasue we can have negative zeroes...
	if r == -0 {
		r = 0
	}
	return complex(r, i)
}
