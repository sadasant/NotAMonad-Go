package monads

// Parses a value out of the given text, and returns the value as well as
// any remaining text.
type Parser func(interface{}) interface{}

// Returns a parser that outputs the given value and consumes no text.
func (p Parser) Wrap(v interface{}) Parser {
	return func(vv interface{}) interface{} {
		return []interface{}{v, vv}
	}
}

// Returns a parser that uses the given parser but transforms its
// resulting value with the given transformation.
func (p Parser) Transform(t Parser) Parser {
	return func(v interface{}) interface{} {
		mid := p(v).([]interface{})
		res := t(mid[0])
		return []interface{}{res, mid[1]}
	}
}

// Returns a parser that uses the given parser to parse an intermediate
// parser, then immediately applies that intermediate parser to the
// remaining text, then returns the resulting value and final remaining
// text.
func (p Parser) Flatten() Parser {
	return func(v interface{}) interface{} {
		mid := p(v).([]interface{})
		q := mid[0].(Parser)
		return q(mid[1])
	}
}
