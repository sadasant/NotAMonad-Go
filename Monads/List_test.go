package monads

import (
	"reflect"
	"testing"
)

func Test_List_Transform(t *testing.T) {
	list := List{1, 2, 3}
	list = list.Transform(func(v Val) Val {
		i := v.(int)
		return i * i
	})
	expected := List{1, 4, 9}
	if reflect.DeepEqual(list, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, list)
	}
}

func Test_List_Flattern(t *testing.T) {
	list := List{
		[]int{},
		[]int{2, 3, 5, 7},
		[]int{1},
		List{0},
		[]string{"banana"},
	}
	list = list.Flatten()
	expected := List{2, 3, 5, 7, 1, 0, "banana"}
	if reflect.DeepEqual(list, expected) != true {
		t.Errorf("DeepEqual\nexpeted:%s\nreceived:%s", expected, list)
	}
}
