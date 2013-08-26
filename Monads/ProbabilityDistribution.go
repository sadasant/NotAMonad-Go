package monads

import (
	"errors"
)

type ProbabilityDistribution map[string]interface{}

// well-formed distributions must add up to 100%:
func (d ProbabilityDistribution) IsOk() error {
	var sum float64
	for _, v := range d {
		if _v, ok := v.([]interface{}); ok {
			sum += _v[1].(float64)
		} else {
			sum += v.(float64)
		}
	}
	if sum != 1 {
		return errors.New("The ProbabilityDistribution is not Well Formed")
	}
	return nil
}

// Returns a probability distribution where the given value is the single
// 100% likely possibility.
func (p ProbabilityDistribution) Wrap(k string) ProbabilityDistribution {
	return ProbabilityDistribution{k: 1.0}
}

// Returns a probability distribution over the result of drawing an input
// from the given distribution, then running that input through the given // transformation.  //
// When two distinct inputs are merged into the same output by the
// transformation, the probability of the output is the sum of the
// inputs' probabilities.
func (p ProbabilityDistribution) Transform(t func(string) string) ProbabilityDistribution {
	r := ProbabilityDistribution{}
	for k, v := range p {
		trans := t(k)
		if _, ok := r[trans]; ok {
			r[trans] = r[trans].(float64) + v.(float64)
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

// Returns a probability distribution over the result of drawing an
// intermediate distribution from the given distribution of
// distributions, then drawing an item from that intermediate
// distribution.
//
// The output probability of an item is its probability times the
// probability of its distribution.  When an item appears in multiple
// intermediate distributions, the corresponding probabilities are added.
func (p ProbabilityDistribution) Flatten() ProbabilityDistribution {
	r := ProbabilityDistribution{}
	for _, l := range p {
		if _l, ok := l.([]interface{}); ok {
			for k, v := range _l[0].(ProbabilityDistribution) {
				if _, ok := r[k]; ok {
					r[k] = r[k].(float64) + (v.(float64) * _l[1].(float64))
				} else {
					if _v, ok := v.([]interface{}); ok {
						r[k] = _v[1].(float64) * _l[1].(float64)
					} else {
						r[k] = v.(float64) * _l[1].(float64)
					}
				}
			}
		} else {
			return p
		}
	}
	return r
}
