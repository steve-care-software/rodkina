package queries

type query struct {
	token  Token
	inside Inside
	fn     QueryFn
}

func createQuery(
	token Token,
	inside Inside,
	fn QueryFn,
) Query {
	out := query{
		token:  token,
		inside: inside,
		fn:     fn,
	}

	return &out
}

// Token returns the token
func (obj *query) Token() Token {
	return obj.token
}

// Inside returns the inside
func (obj *query) Inside() Inside {
	return obj.inside
}

// Fn returns the func
func (obj *query) Fn() QueryFn {
	return obj.fn
}
