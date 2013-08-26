package nonmonads

import (
	"reflect"
	"testing"
)

func Test_SquareMatrix_Wrap(t *testing.T) {
	s := new(SquareMatrix).Wrap("test")
	if len(s.Items) != 1 {
		t.Errorf("len(s)\nexpeted:%s\nreceived:%s", 1, len(s.Items))
	}
	if s.Items[0][0] != "test" {
		t.Errorf("s[0][0]\nexpeted:%s\nreceived:%s", "test", s.Items[0][0])
	}
}

func Test_SquareMatrix_Transform(t *testing.T) {
	var s, expected SquareMatrix
	s.From(Square{[]interface{}{3, 7}, []interface{}{2, 5}})
	s, _ = s.Transform(func(v interface{}) interface{} {
		i := v.(int)
		return i * i
	})
	expected.From(Square{[]interface{}{9, 49}, []interface{}{4, 25}})
	if reflect.DeepEqual(s, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, s)
	}
}

func Test_SquareMatrix_Flatten(t *testing.T) {
	var s1, s2, r3, s3, r4, s4, s5 SquareMatrix
	var expected SquareMatrix
	s1, _ = new(SquareMatrix).Wrap(new(SquareMatrix).Wrap("test")).Flatten()
	if len(s1.Items) != 1 {
		t.Errorf("len(s1)expeted:%s\nreceived:%s", 1, len(s1.Items))
	}
	if s1.Items[0][0] != "test" {
		t.Errorf("s1[0][0]\nexpeted:%s\nreceived:%s", "test", s1.Items[0][0])
	}
	s2, _ = SquareMatrix{}.Flatten()
	if len(s2.Items) != 0 {
		t.Errorf("len(s2)expeted:%s\nreceived:%s", 1, len(s2.Items))
	}
	r3.From(Square{[]interface{}{}, []interface{}{}})
	r3.Items[0] = append(r3.Items[0], SquareMatrix{})
	s3, _ = r3.Flatten()
	if len(s3.Items) != 0 {
		t.Errorf("len(s2)expeted:%s\nreceived:%s", 0, len(s3.Items))
	}
	r4.From(Square{[]interface{}{3, 7}, []interface{}{2, 5}})
	s4, _ = r4.Transform(func(v interface{}) interface{} {
		return new(SquareMatrix).Wrap(v)
	})
	s4, _ = s4.Flatten()
	expected.From(Square{[]interface{}{3, 7}, []interface{}{2, 5}})
	if reflect.DeepEqual(s4.Items, expected.Items) != true {
		t.Errorf("DeepEqual\nexpeted: %s\nreceived:%s", expected.Items, s4.Items)
	}
	s5, _ = r4.Transform(func(v interface{}) interface{} {
		return s4
	})
	s5, _ = s5.Flatten()
	expected.From(Square{
		[]interface{}{3, 7, 3, 7},
		[]interface{}{2, 5, 2, 5},
		[]interface{}{3, 7, 3, 7},
		[]interface{}{2, 5, 2, 5},
	})
	if reflect.DeepEqual(s5.Items, expected.Items) != true {
		t.Errorf("DeepEqual\nexpeted: %s\nreceived:%s", expected.Items, s5.Items)
	}
}

func Test_SquareMatrix_Flatten_Broken(t *testing.T) {
	var r, bad SquareMatrix
	bad.From(Square{})
	r.From(Square{
		[]interface{}{
			new(SquareMatrix).Wrap(1),
			bad,
		},
		[]interface{}{
			new(SquareMatrix).Wrap(1),
			new(SquareMatrix).Wrap(1),
		},
	})
	s, err := r.Flatten()
	// t.Errorf("%s", s.Items)
	if err == nil {
		t.Error("SquareMatrix should not be well formed: %s", s.Items)
	}
}
