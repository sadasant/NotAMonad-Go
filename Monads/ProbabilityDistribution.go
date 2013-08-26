package monads

import (
	"errors"
)

type Distribution map[string]interface{}

// PD instead of ProbabilityDistribution
type PD func() Distribution

func (p PD) From(dict Distribution) (PD, error) {
	return func() Distribution {
		return dict
	}, dict.isWellFormed()
}

// well-formed distributions must add up to 100%:
func (d Distribution) isWellFormed() error {
	var sum float64
	for _, v := range d {
		if _v, ok := v.([]interface{}); ok {
			sum += _v[1].(float64)
		} else {
			sum += v.(float64)
		}
	}
	if sum != 1 {
		return errors.New("The PD is not Well Formed")
	}
	return nil
}

// Returns a probability distribution where the given value is the single
// 100% likely possibility.
func (p PD) Wrap(k string) PD {
	p, _ = p.From(Distribution{k: 1.0})
	return p
}

// Returns a probability distribution over the result of drawing an input
// from the given distribution, then running that input through the given // transformation.  //
// When two distinct inputs are merged into the same output by the
// transformation, the probability of the output is the sum of the
// inputs' probabilities.
func (p PD) Transform(t func(string) string) (PD, error) {
	pd := p()
	rd := Distribution{}
	for k, v := range pd {
		trans := t(k)
		if _, ok := rd[trans]; ok {
			rd[trans] = rd[trans].(float64) + v.(float64)
		} else {
			if _v, ok := v.([]interface{}); ok {
				rd[trans] = _v[1]
			} else {
				rd[trans] = v
			}
		}
	}
	return p.From(rd)
}

// Returns a probability distribution over the result of drawing an
// intermediate distribution from the given distribution of
// distributions, then drawing an item from that intermediate
// distribution.
//
// The output probability of an item is its probability times the
// probability of its distribution.  When an item appears in multiple
// intermediate distributions, the corresponding probabilities are added.
func (p PD) Flatten() (PD, error) {
	pd := p()
	rd := Distribution{}
	for _, l := range pd {
		if _l, ok := l.([]interface{}); ok {
			for k, v := range _l[0].(PD)() {
				if _, ok := rd[k]; ok {
					rd[k] = rd[k].(float64) + (v.(float64) * _l[1].(float64))
				} else {
					if _v, ok := v.([]interface{}); ok {
						rd[k] = _v[1].(float64) * _l[1].(float64)
					} else {
						rd[k] = v.(float64) * _l[1].(float64)
					}
				}
			}
		} else {
			return p, nil
		}
	}
	return p.From(rd)
}
