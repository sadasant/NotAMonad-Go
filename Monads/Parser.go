package monads

// Parses a value out of the given text, and returns the value as well as
// any remaining text.
type Parser func(Val) Val

func (p Parser) Wrap(v Val) Parser {
	return func(vv Val) Val {
		return List{v, vv}
	}
}

// Returns a parser that uses the given parser but transforms its
// resulting value with the given transformation.
func (p Parser) Transform(t Parser) Parser {
	return func(v Val) Val {
		mid := p(v).(List)
		res := t(mid[0])
		return List{res, mid[1]}
	}
}

// Returns a parser that uses the given parser to parse an intermediate
// parser, then immediately applies that intermediate parser to the
// remaining text, then returns the resulting value and final remaining
// text.
func (p Parser) Flatten() Parser {
	return func(v Val) Val {
		mid := p(v).(List)
		q := mid[0].(Parser)
		return q(mid[1])
	}
}
