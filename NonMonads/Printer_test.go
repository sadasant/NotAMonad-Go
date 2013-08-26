package nonmonads

import (
	"reflect"
	"strconv"
	"testing"
)

func Example_Printer_Wrap() {
	s := new(Printer).Wrap("test")
	s("input")
	s("test")
	s("a")
	// Output:
	// test
	// test
	// test
}

func Test_Printer_Transform(t *testing.T) {
	var r []int
	var p Printer
	p = func(m interface{}) {
		i, _ := strconv.Atoi(m.(string))
		r = append(r, i)
	}
	q := p.Transform(func(m interface{}) interface{} {
		return strconv.Itoa(m.(int))
	})
	q(5)
	expected := []int{5}
	if reflect.DeepEqual(r, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, r)
	}
	q(8)
	expected = []int{5, 8}
	if reflect.DeepEqual(r, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, r)
	}
}

func Test_Printer_Flatten(t *testing.T) {
	var r []interface{}
	var p Printer
	p = func(m interface{}) {
		r = append(r, m)
	}
	q := p.Flatten()
	// AWKWARD: can't really verify what the printers are, just that they
	// were made and inserted
	q(5)
	if len(r) != 1 {
		t.Errorf("len(r)\nexpeted:%s\nreceived:%s", 1, len(r))
	}
	q(8)
	if len(r) != 2 {
		t.Errorf("len(r)\nexpeted:%s\nreceived:%s", 2, len(r))
	}
}
