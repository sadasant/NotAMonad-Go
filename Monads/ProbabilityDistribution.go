package monads

import (
	"errors"
)

type Distribution map[string]interface{}

type ProbabilityDistribution struct {
	Dict Distribution
}

// Returns a probability distribution where the given value is the single
// 100% likely possibility.
func (p *ProbabilityDistribution) Wrap(k string) ProbabilityDistribution{
	p.Dict = Distribution{k: 1.0}
	return *p
}

func (p *ProbabilityDistribution) From(dict map[string]interface{}) error {
	p.Dict = dict
	return p.isWellFormed()
}

// well-formed distributions must add up to 100%:
func (p ProbabilityDistribution) isWellFormed() error {
	var sum float64
	for _, v := range p.Dict {
		if _v, ok := v.(List); ok {
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

// Returns a probability distribution over the result of drawing an input
// from the given distribution, then running that input through the given // transformation.  //
// When two distinct inputs are merged into the same output by the
// transformation, the probability of the output is the sum of the
// inputs' probabilities.
func (p ProbabilityDistribution) Transform(t func(string) string) (r ProbabilityDistribution) {
	r.Dict = map[string]interface{}{}
	for k, v := range p.Dict {
		trans := t(k)
		if _, ok := r.Dict[trans]; ok {
			r.Dict[trans] = r.Dict[trans].(float64) + v.(float64)
		} else {
			if _v, ok := v.(List); ok {
				r.Dict[trans] = _v[1]
			} else {
				r.Dict[trans] = v
			}
		}
	}
	return
}

// Returns a probability distribution over the result of drawing an
// intermediate distribution from the given distribution of
// distributions, then drawing an item from that intermediate
// distribution.
//
// The output probability of an item is its probability times the
// probability of its distribution.  When an item appears in multiple
// intermediate distributions, the corresponding probabilities are added.
func (p ProbabilityDistribution) Flatten() (r ProbabilityDistribution) {
	r.Dict = map[string]interface{}{}
	for _, l := range p.Dict {
		if _l, ok := l.(List); ok {
			for k, v := range _l[0].(ProbabilityDistribution).Dict {
				if _, ok := r.Dict[k]; ok {
					r.Dict[k] = r.Dict[k].(float64) + (v.(float64) * _l[1].(float64))
				} else {
					if _v, ok := v.(List); ok {
						r.Dict[k] = _v[1].(float64) * _l[1].(float64)
					} else {
						r.Dict[k] = v.(float64) * _l[1].(float64)
					}
				}
			}
		} else {
			return p
		}
	}
	return
}
