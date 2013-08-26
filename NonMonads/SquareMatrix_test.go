package nonmonads

import (
	"reflect"
	"testing"
)

func Test_SquareMatrix_Wrap(t *testing.T) {
	s := new(SquareMatrix).Wrap("test")
	if len(s) != 1 {
		t.Errorf("len(s)\nexpeted:%s\nreceived:%s", 1, len(s))
	}
	if s[0][0] != "test" {
		t.Errorf("s[0][0]\nexpeted:%s\nreceived:%s", "test", s[0][0])
	}
}

func Test_SquareMatrix_Transform(t *testing.T) {
	r := SquareMatrix{[]interface{}{3, 7}, []interface{}{2, 5}}
	s := r.Transform(func(v interface{}) interface{} {
		i := v.(int)
		return i * i
	})
	expected := SquareMatrix{[]interface{}{9, 49}, []interface{}{4, 25}}
	if reflect.DeepEqual(s, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, s)
	}
}

func Test_SquareMatrix_Flatten(t *testing.T) {
	var expected SquareMatrix
	s1 := new(SquareMatrix).Wrap(new(SquareMatrix).Wrap("test")).Flatten()
	if len(s1) != 1 {
		t.Errorf("len(s1)expeted:%s\nreceived:%s", 1, len(s1))
	}
	if s1[0][0] != "test" {
		t.Errorf("s1[0][0]\nexpeted:%s\nreceived:%s", "test", s1[0][0])
	}
	s2 := SquareMatrix{}.Flatten()
	if len(s2) != 0 {
		t.Errorf("len(s2)expeted:%s\nreceived:%s", 1, len(s2))
	}
	r3 := SquareMatrix{[]interface{}{}, []interface{}{}}
	r3[0] = append(r3[0], SquareMatrix{})
	s3 := r3.Flatten()
	if len(s3) != 0 {
		t.Errorf("len(s2)expeted:%s\nreceived:%s", 0, len(s3))
	}
	r4 := SquareMatrix{[]interface{}{3, 7}, []interface{}{2, 5}}
	s4 := r4.Transform(func(v interface{}) interface{} {
		return new(SquareMatrix).Wrap(v)
	})
	s4 = s4.Flatten()
	expected = SquareMatrix{[]interface{}{3, 7}, []interface{}{2, 5}}
	if reflect.DeepEqual(s4, expected) != true {
		t.Errorf("DeepEqual\nexpeted: %s\nreceived:%s", expected, s4)
	}
	s5 := r4.Transform(func(v interface{}) interface{} {
		return s4
	})
	s5 = s5.Flatten()
	expected = SquareMatrix{
		[]interface{}{3, 7, 3, 7},
		[]interface{}{2, 5, 2, 5},
		[]interface{}{3, 7, 3, 7},
		[]interface{}{2, 5, 2, 5},
	}
	if reflect.DeepEqual(s5, expected) != true {
		t.Errorf("DeepEqual\nexpeted: %s\nreceived:%s", expected, s5)
	}
}

func Test_SquareMatrix_Flatten_Broken(t *testing.T) {
	var bad SquareMatrix
	r := SquareMatrix{
		[]interface{}{
			new(SquareMatrix).Wrap(1),
			bad,
		},
		[]interface{}{
			new(SquareMatrix).Wrap(1),
			new(SquareMatrix).Wrap(1),
		},
	}
	s := r.Flatten()
	err := s.IsOk()
	// t.Errorf("%s", s)
	if err == nil {
		t.Error("SquareMatrix should not be well formed: %s", s)
	}
}
