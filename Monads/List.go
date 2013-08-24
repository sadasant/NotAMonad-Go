package monads

import (
// 	"reflect"
)

type Val interface{}
type List []Val

// Returns a list created by running all of the items in the given list
// through a transformation function.
func (l List) Transform(f func(Val) Val) (t List) {
	for _, v := range l {
		t = append(t, f(v))
	}
	return
}

func (l List) Flatten() (t List) {
	for _, v := range l {
		switch v.(type) {
		case List:
			for _, vv := range v.(List) {
				t = append(t, vv)
			}
		case []Val:
			for _, vv := range v.([]Val) {
				t = append(t, vv)
			}
		case []string:
			for _, vv := range v.([]string) {
				t = append(t, vv)
			}
		case []int:
			for _, vv := range v.([]int) {
				t = append(t, vv)
			}
		default:
			t = append(t, v)
		}
	}
	return
}
