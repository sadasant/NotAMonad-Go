package monads

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_Parser_Wrap(t *testing.T) {
	list := new(Parser).Wrap("test")(nil)
	expected := []interface{}{"test", nil}
	if reflect.DeepEqual(list, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, list)
	}
}

func Test_Parser_Transform(t *testing.T) {
	p := Parser(func(v interface{}) interface{} {
		s := v.(string)
		return []interface{}{s[0:5], s[5:]}
	})
	q := p.Transform(func(v interface{}) interface{} {
		i, _ := strconv.Atoi(v.(string))
		return i * i
	})
	expected := []interface{}{25, "wonder"}
	received := q("00005wonder")
	if reflect.DeepEqual(received, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, received)
	}
}

func Test_Parser_Flatten(t *testing.T) {
	psubs := Parser(func(v interface{}) interface{} {
		s := v.(string)
		i, _ := strconv.Atoi(s[0:5])
		return []interface{}{i, s[5:]}
	})
	p := Parser(func(v interface{}) interface{} {
		s := v.(string)
		if s[0] == 'a' {
			return []interface{}{psubs, s[1:]}
		}
		return []interface{}{psubs.Wrap(-1), s[1:]}
	})
	q := p.Flatten()
	expected := []interface{}{4, "x"}
	received := q("a00004x")
	if reflect.DeepEqual(received, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, received)
	}
	expected = []interface{}{-1, "00004x"}
	received = q("b00004x")
	if reflect.DeepEqual(received, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, received)
	}
}
