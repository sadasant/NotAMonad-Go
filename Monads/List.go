package monads

type List []interface{}

// Returns a list containing a single value: the given value.
func (l *List) Wrap(v interface{}) List {
    return append(*l, v)
}

// Returns a list created by running all of the items in the given list
// through a transformation function.
func (l List) Transform(t func(interface{}) interface{}) (r List) {
	for _, v := range l {
		r = append(r, t(v))
	}
	return
}

// Returns a list created by appending together all of the items in all
// of the lists in the given list.
func (l List) Flatten() (r List) {
	for _, v := range l {
		switch v.(type) {
		case List:
			for _, vv := range v.(List) {
				r = append(r, vv)
			}
		case []interface{}:
			for _, vv := range v.([]interface{}) {
				r = append(r, vv)
			}
		case []string:
			for _, vv := range v.([]string) {
				r = append(r, vv)
			}
		case []int:
			for _, vv := range v.([]int) {
				r = append(r, vv)
			}
		default:
			r = append(r, v)
		}
	}
	return
}
