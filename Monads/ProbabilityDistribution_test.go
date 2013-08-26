package monads

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_PD_Wrap(t *testing.T) {
	p := new(PD).Wrap("test")
	expected := Distribution{"test": 1.0}
	if reflect.DeepEqual(p(), expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, p())
	}
}

func Test_PD_Transform(t *testing.T) {
	p := PD(func() Distribution {
		return Distribution{
			"hey":    1.0 / 2,
			"listen": 1.0 / 4,
			"die":    1.0 / 4,
		}
	})
	err := p().IsOk()
	if err != nil {
		t.Errorf("PD %s caused: %s", p(), err)
	}
	q := p.Transform(func(k string) string {
		return strconv.Itoa(len(k))
	})
	expected := Distribution{
		"3": 3.0 / 4,
		"6": 1.0 / 4,
	}
	if reflect.DeepEqual(q(), expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, q())
	}
}

func Test_PD_Flatten(t *testing.T) {
	p := new(PD).Wrap("test").Flatten()
	expected := Distribution{"test": 1.0}
	if reflect.DeepEqual(p(), expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, p())
	}

	a1 := PD(func() Distribution {
		return Distribution{
			"a":   1.0 / 7,
			"bra": 1.0 / 7 * 2,
			"ca":  1.0 / 7 * 4,
		}
	})
	a2 := PD(func() Distribution {
		return Distribution{
			"da":  1.0 / 3,
			"bra": 1.0 / 3 * 2,
		}
	})
	s2 := PD(func() Distribution {
		return Distribution{
			"a1": []interface{}{a1, 1.0 / 5},
			"a2": []interface{}{a2, 1.0 / 5 * 4},
		}
	})
	s2 = s2.Flatten()
	expected = Distribution{
		"a":   1.0 / 105 * 3,
		"bra": 1.0 / 105 * 62,
		"ca":  1.0 / 105 * 12,
		"da":  1.0 / 105 * 28,
	}
	if reflect.DeepEqual(s2(), expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, s2())
	}
}
