package nonmonads

import (
	"math"
	"reflect"
	"strconv"
	"testing"
)

func Test_QS_Wrap(t *testing.T) {
	s := new(QS).Wrap("test")
	expected := Superposition{"test": complex(1.0, 0)}
	if reflect.DeepEqual(s(), expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, s())
	}
}

func Test_QS_Transform(t *testing.T) {
	s := QS(func() Superposition {
		return Superposition{
			"hey": complex(1.0/5*3, 0),
			"you": -complex(1.0/5*4, 0),
		}
	})
	s = s.Transform(func(k string) string {
		return strconv.Itoa(len(k))
	})
	err := s().IsOk()
	if err == nil {
		t.Errorf("QS %s shouldn't be well construct", s())
	}
}

func Test_QS_Broken(t *testing.T) {
	s := QS(func() Superposition {
		return Superposition{
			"hey":    complex(1.0/5*3, 0),
			"listen": complex(1.0/5*4, 0),
		}
	})
	s = s.Transform(func(k string) string {
		return strconv.Itoa(len(k))
	})
	expected := Superposition{
		"3": complex(1.0/5*3, 0),
		"6": complex(1.0/5*4, 0),
	}
	if reflect.DeepEqual(s(), expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, s())
	}
}

func Test_QS_Flatten(t *testing.T) {
	a1 := QS(func() Superposition {
		return Superposition{
			"hey":    complex(1.0/5*3, 0),
			"listen": complex(-1.0/5*4, 0),
		}
	})
	a2 := QS(func() Superposition {
		return Superposition{
			"over":  complex(1.0/5*3, 0),
			"there": complex(-1.0/5*4, 0),
		}
	})
	s2 := QS(func() Superposition {
		return Superposition{
			"a1": []interface{}{a1, complex(1.0/5*3, 0)},
			"a2": []interface{}{a2, complex(0, 1.0/5*4)},
		}
	})
	s2 = s2.Flatten()
	expected := Superposition{
		"hey":    complex(1.0/25*9, 0),
		"listen": complex(-1.0/25*12, 0),
		"over":   complex(0, 1.0/25*12),
		"there":  complex(0, -1.0/25*16),
	}
	// Fix rounding.
	c := s2()["there"].(complex128)
	i := (math.Floor(imag(c)*100) / 100) + 0.01
	s2()["there"] = complex(real(c), i)
	if reflect.DeepEqual(s2(), expected) != true {
		t.Errorf("DeepEqual\nexpeted: %s\nreceived:%s", expected, s2())
	}
}

func Test_QS_Flatten_Broken(t *testing.T) {
	a1 := QS(func() Superposition {
		return Superposition{"hey": complex(1.0, 0)}
	})
	a2 := QS(func() Superposition {
		return Superposition{"hey": complex(-1.0, 0)}
	})
	s2 := QS(func() Superposition {
		return Superposition{
			"a1": []interface{}{a1, complex(1.0/5*3, 0)},
			"a2": []interface{}{a2, complex(1.0/5*4, 0)},
		}
	})
	s2 = s2.Flatten()
	err := s2().IsOk()
	if err == nil {
		t.Errorf("QS %s shouldn't be well construct", s2())
	}
}
