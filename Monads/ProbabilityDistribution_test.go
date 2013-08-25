package monads

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_ProbabilityDistribution_Transform(t *testing.T) {
	var p ProbabilityDistribution
	err := p.From(Distribution{
		"hey":    1.0 / 2,
		"listen": 1.0 / 4,
		"die":    1.0 / 4,
	})
	if err != nil {
		t.Errorf("ProbabilityDistribution %s caused: %s", p.Dict, err)
	}
	q := p.Transform(func(k string) string {
		return strconv.Itoa(len(k))
	})
	expected := Distribution{
		"3": 3.0 / 4,
		"6": 1.0 / 4,
	}
	if reflect.DeepEqual(q.Dict, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, q.Dict)
	}
}

func Test_ProbabilityDistribution_Flatten(t *testing.T) {
	var p ProbabilityDistribution
	err := p.From(Distribution{"test": 1.0})
	p = p.Flatten()
	expected := Distribution{"test": 1.0}
	if reflect.DeepEqual(p.Dict, expected) != true {
		t.Errorf("ProbabilityDistribution %s caused: %s", p.Dict, err)
	}

	var a1, a2, s2 ProbabilityDistribution
	a1.From(Distribution{
		"a":   1.0 / 7,
		"bra": 1.0 / 7 * 2,
		"ca":  1.0 / 7 * 4,
	})
	a2.From(Distribution{
		"da":  1.0 / 3,
		"bra": 1.0 / 3 * 2,
	})
	s2.From(Distribution{
		"a1": List{a1, 1.0 / 5},
		"a2": List{a2, 1.0 / 5 * 4},
	})
	s2 = s2.Flatten()
	expected = Distribution{
		"a":   1.0 / 105 * 3,
		"bra": 1.0 / 105 * 62,
		"ca":  1.0 / 105 * 12,
		"da":  1.0 / 105 * 28,
	}
	if reflect.DeepEqual(s2.Dict, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, s2.Dict)
	}
}
