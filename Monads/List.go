package monads

type Val interface{}
type List []Val

// Returns a list containing a single value: the given value.
func (l *List) Wrap(v Val) List {
    return append(*l, v)
}

// Returns a list created by running all of the items in the given list
// through a transformation function.
func (l List) Transform(t func(Val) Val) (r List) {
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
		case []Val:
			for _, vv := range v.([]Val) {
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
