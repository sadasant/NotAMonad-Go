package monads

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_ProbabilityDistribution_Wrap(t *testing.T) {
	p := new(ProbabilityDistribution).Wrap("test")
	expected := ProbabilityDistribution{"test": 1.0}
	if reflect.DeepEqual(p, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, p)
	}
}

func Test_ProbabilityDistribution_Transform(t *testing.T) {
	p := ProbabilityDistribution{
		"hey":    1.0 / 2,
		"listen": 1.0 / 4,
		"die":    1.0 / 4,
	}
	err := p.IsOk()
	if err != nil {
		t.Errorf("ProbabilityDistribution %s caused: %s", p, err)
	}
	q := p.Transform(func(k string) string {
		return strconv.Itoa(len(k))
	})
	expected := ProbabilityDistribution{
		"3": 3.0 / 4,
		"6": 1.0 / 4,
	}
	if reflect.DeepEqual(q, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, q)
	}
}

func Test_ProbabilityDistribution_Flatten(t *testing.T) {
	p := new(ProbabilityDistribution).Wrap("test").Flatten()
	expected := ProbabilityDistribution{"test": 1.0}
	if reflect.DeepEqual(p, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, p)
	}

	a1 := ProbabilityDistribution{
		"a":   1.0 / 7,
		"bra": 1.0 / 7 * 2,
		"ca":  1.0 / 7 * 4,
	}
	a2 := ProbabilityDistribution{
		"da":  1.0 / 3,
		"bra": 1.0 / 3 * 2,
	}
	s2 := ProbabilityDistribution{
		"a1": []interface{}{a1, 1.0 / 5},
		"a2": []interface{}{a2, 1.0 / 5 * 4},
	}

	s2 = s2.Flatten()
	expected = ProbabilityDistribution{
		"a":   1.0 / 105 * 3,
		"bra": 1.0 / 105 * 62,
		"ca":  1.0 / 105 * 12,
		"da":  1.0 / 105 * 28,
	}
	if reflect.DeepEqual(s2, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, s2)
	}
}
