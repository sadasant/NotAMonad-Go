package nonmonads

import (
	"errors"
	"math/cmplx"
)

type QuantumSuperposition map[string]interface{}

// well-formed superpositions must add up to 100%:
func (p QuantumSuperposition) IsOk() error {
	var sum float64
	for _, v := range p {
		if _v, ok := v.([]interface{}); ok {
			sum += cmplx.Abs(_v[1].(complex128))
		} else {
			sum += cmplx.Abs(v.(complex128))
		}
	}
	if sum != 1 {
		return errors.New("The QuantumSuperposition is not Well Formed")
	}
	return nil
}

// Returns a superposition where the given value is the single state and
// has an amplitude of 1.
func (q QuantumSuperposition) Wrap(k string) QuantumSuperposition {
	return QuantumSuperposition{k: complex(1.0, 0)}
}

// Returns a superposition over the result of drawing an input from the
// given superposition, then running that input through the given
// transformation.
//
// BROKEN - When distinct inputs are merged, they interfere.
// The interference breaks the squared magnitude constraint.
func (q QuantumSuperposition) Transform(t func(string) string) QuantumSuperposition {
	r := QuantumSuperposition{}
	for k, v := range q {
		trans := t(k)
		if _, ok := r[trans]; ok {
			r[trans] = r[trans].(complex128) + v.(complex128)
		} else {
			if _v, ok := v.([]interface{}); ok {
				r[trans] = _v[1]
			} else {
				r[trans] = v
			}
		}
	}
	return r
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
func (q QuantumSuperposition) Flatten() QuantumSuperposition {
	r := QuantumSuperposition{}
	for _, l := range q {
		if _l, ok := l.([]interface{}); ok {
			for k, v := range _l[0].(QuantumSuperposition) {
				if _, ok := r[k]; ok {
					r[k] = r[k].(complex128) + mult(v.(complex128), _l[1].(complex128))
				} else {
					if _v, ok := v.([]interface{}); ok {
						r[k] = mult(_v[1].(complex128), _l[1].(complex128))
					} else {
						r[k] = mult(v.(complex128), _l[1].(complex128))
					}
				}
			}
		} else {
			return q
		}
	}
	return r
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
