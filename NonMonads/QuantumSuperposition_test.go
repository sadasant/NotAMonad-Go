package nonmonads

import (
	"math"
	"reflect"
	"strconv"
	"testing"
)

func Test_QuantumSuperposition_Wrap(t *testing.T) {
	s := new(QuantumSuperposition).Wrap("test")
	expected := Superposition{"test": complex(1.0, 0)}
	if reflect.DeepEqual(s.Dict, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, s.Dict)
	}
}

func Test_QuantumSuperposition_Transform(t *testing.T) {
	s, _ := new(QuantumSuperposition).From(Superposition{
		"hey": complex(1.0/5*3, 0),
		"you": -complex(1.0/5*4, 0),
	})
	_, err := s.Transform(func(k string) string {
		return strconv.Itoa(len(k))
	})
	if err == nil {
		t.Errorf("QuantumSuperposition %s shouldn't be well construct", s.Dict)
	}
}

func Test_QuantumSuperposition_Broken(t *testing.T) {
	s, _ := new(QuantumSuperposition).From(Superposition{
		"hey":    complex(1.0/5*3, 0),
		"listen": complex(1.0/5*4, 0),
	})
	s, _ = s.Transform(func(k string) string {
		return strconv.Itoa(len(k))
	})
	expected := Superposition{
		"3": complex(1.0/5*3, 0),
		"6": complex(1.0/5*4, 0),
	}
	if reflect.DeepEqual(s.Dict, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, s.Dict)
	}
}

func Test_QuantumSuperposition_Flatten(t *testing.T) {
	a1, _ := new(QuantumSuperposition).From(Superposition{
		"hey":    complex(1.0/5*3, 0),
		"listen": complex(-1.0/5*4, 0),
	})
	a2, _ := new(QuantumSuperposition).From(Superposition{
		"over":  complex(1.0/5*3, 0),
		"there": complex(-1.0/5*4, 0),
	})
	s2, _ := new(QuantumSuperposition).From(Superposition{
		"a1": []Val{a1, complex(1.0/5*3, 0)},
		"a2": []Val{a2, complex(0, 1.0/5*4)},
	})
	s2, _ = s2.Flatten()
	expected := Superposition{
		"hey":    complex(1.0/25*9, 0),
		"listen": complex(-1.0/25*12, 0),
		"over":   complex(0, 1.0/25*12),
		"there":  complex(0, -1.0/25*16),
	}
	// Fix rounding.
	c := s2.Dict["there"].(complex128)
	i := (math.Floor(imag(c)*100) / 100) + 0.01
	s2.Dict["there"] = complex(real(c), i)
	if reflect.DeepEqual(s2.Dict, expected) != true {
		t.Errorf("DeepEqual\nexpeted: %s\nreceived:%s", expected, s2.Dict)
	}
}

func Test_QuantumSuperposition_Flatten_Broken(t *testing.T) {
	a1, _ := new(QuantumSuperposition).From(Superposition{
		"hey": complex(1.0, 0),
	})
	a2, _ := new(QuantumSuperposition).From(Superposition{
		"hey": complex(-1.0, 0),
	})
	s2, _ := new(QuantumSuperposition).From(Superposition{
		"a1": []Val{a1, complex(1.0/5*3, 0)},
		"a2": []Val{a2, complex(1.0/5*4, 0)},
	})
	s2, err := s2.Flatten()
	if err == nil {
		t.Errorf("QuantumSuperposition %s shouldn't be well construct", s2.Dict)
	}
}
