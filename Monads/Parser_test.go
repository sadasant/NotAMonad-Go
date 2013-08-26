package monads

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_Parser_Wrap(t *testing.T) {
	list := new(Parser).Wrap("test")(nil)
	expected := List{"test", nil}
	if reflect.DeepEqual(list, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, list)
	}
}

func Test_Parser_Transform(t *testing.T) {
	p := Parser(func(v interface{}) interface{} {
		s := v.(string)
		return List{s[0:5], s[5:]}
	})
	q := p.Transform(func(v interface{}) interface{} {
		i, _ := strconv.Atoi(v.(string))
		return i * i
	})
	expected := List{25, "wonder"}
	received := q("00005wonder")
	if reflect.DeepEqual(received, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, received)
	}
}

func Test_Parser_Flatten(t *testing.T) {
	psubs := Parser(func(v interface{}) interface{} {
		s := v.(string)
		i, _ := strconv.Atoi(s[0:5])
		return List{i, s[5:]}
	})
	p := Parser(func(v interface{}) interface{} {
		s := v.(string)
		if s[0] == 'a' {
			return List{psubs, s[1:]}
		}
		return List{psubs.Wrap(-1), s[1:]}
	})
	q := p.Flatten()
	expected := List{4, "x"}
	received := q("a00004x")
	if reflect.DeepEqual(received, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, received)
	}
	expected = List{-1, "00004x"}
	received = q("b00004x")
	if reflect.DeepEqual(received, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, received)
	}
}
