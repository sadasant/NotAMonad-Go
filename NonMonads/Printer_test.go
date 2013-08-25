package nonmonads

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_Printer_Wrap(t *testing.T) {
	s := new(Printer).Wrap("test")
	expected := "test"
	if received := s("input"); received != expected {
		t.Errorf("Printer\nexpeted:%s\nreceived:%s", expected, received)
	}
	if received := s("test"); received != expected {
		t.Errorf("Printer\nexpeted:%s\nreceived:%s", expected, received)
	}
	if received := s("a"); received != expected {
		t.Errorf("Printer\nexpeted:%s\nreceived:%s", expected, received)
	}
}

func Test_Printer_Transform(t *testing.T) {
	var r []int
	var p Printer
	p = func(m Val) Val {
		i, _ := strconv.Atoi(m.(string))
		r = append(r, i)
		return m
	}
	q := p.Transform(func(m Val) Val {
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
	var r []int
	var p Printer
	p = func(m Val) Val {
		r = append(r, m.(Printer)("").(int))
		return m
	}
	q := p.Flatten()
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
