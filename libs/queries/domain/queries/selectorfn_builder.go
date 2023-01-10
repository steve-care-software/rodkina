package queries

import "errors"

type queryFnBuilder struct {
	single SingleQueryFn
	multi  MultiQueryFn
}

func createQueryFnBuilder() QueryFnBuilder {
	out := queryFnBuilder{
		single: nil,
		multi:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *queryFnBuilder) Create() QueryFnBuilder {
	return createQueryFnBuilder()
}

// WithSingle adds a single func to the builder
func (app *queryFnBuilder) WithSingle(single SingleQueryFn) QueryFnBuilder {
	app.single = single
	return app
}

// WithMulti adds a multi func to the builder
func (app *queryFnBuilder) WithMulti(multi MultiQueryFn) QueryFnBuilder {
	app.multi = multi
	return app
}

// Now builds a new Query func
func (app *queryFnBuilder) Now() (QueryFn, error) {
	if app.single != nil {
		return createQueryFnWithSingle(app.single), nil
	}

	if app.multi != nil {
		return createQueryFnWithMulti(app.multi), nil
	}

	return nil, errors.New("the QueryFn is invalid")
}
